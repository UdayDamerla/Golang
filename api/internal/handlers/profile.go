package handlers

import (
	"net/http"

	"api/internal/middleware"
	"api/pkg/response"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Context().Value(middleware.ContextUserKey).(string)
	response.JSON(w, http.StatusOK, response.APIResponse{Data: map[string]string{"username": username}})
}
