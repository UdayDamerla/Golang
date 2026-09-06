package handlers

import (
	"encoding/json"
	"net/http"

	"api/internal/services"
	"api/pkg/response"
)

type AuthHandler struct {
	auth *services.AuthService
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAuthHandler(auth *services.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Error: "invalid request body"})
		return
	}

	if err := h.auth.Register(creds.Username, creds.Password); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Error: err.Error()})
		return
	}

	response.JSON(w, http.StatusCreated, response.APIResponse{Message: "user registered"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		response.JSON(w, http.StatusBadRequest, response.APIResponse{Error: "invalid request body"})
		return
	}

	token, err := h.auth.Login(creds.Username, creds.Password)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, response.APIResponse{Error: err.Error()})
		return
	}

	response.JSON(w, http.StatusOK, response.APIResponse{Message: "login successful", Data: map[string]string{"token": token}})
}
