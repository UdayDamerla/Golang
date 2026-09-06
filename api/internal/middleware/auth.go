package middleware

import (
	"context"
	"net/http"
	"strings"

	"api/internal/services"
	"api/pkg/response"
)

type contextKey string

const ContextUserKey contextKey = "user"

func JWTAuth(authService *services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				response.JSON(w, http.StatusUnauthorized, response.APIResponse{Error: "missing or invalid authorization header"})
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := authService.ValidateToken(token)
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, response.APIResponse{Error: "invalid or expired token"})
				return
			}

			ctx := context.WithValue(r.Context(), ContextUserKey, claims.Username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
