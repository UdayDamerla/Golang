package services

import (
	"errors"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	PasswordHash string
}

type AuthService struct {
	jwtSecret []byte
	users     map[string]User
	mu        sync.RWMutex
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{
		jwtSecret: []byte(secret),
		users:     make(map[string]User),
	}
}

func (a *AuthService) Register(username, password string) error {
	if username == "" || password == "" {
		return errors.New("username and password are required")
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	if _, exists := a.users[username]; exists {
		return errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	a.users[username] = User{
		Username:     username,
		PasswordHash: string(hash),
	}
	return nil
}

func (a *AuthService) Login(username, password string) (string, error) {
	a.mu.RLock()
	user, exists := a.users[username]
	a.mu.RUnlock()

	if !exists {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return a.generateToken(username)
}

func (a *AuthService) generateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(a.jwtSecret)
	if err != nil {
		return "", errors.New("failed to sign token")
	}

	return signedToken, nil
}

func (a *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return a.jwtSecret, nil
	})
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
