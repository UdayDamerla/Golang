package handlers

import (
	"net/http"

	"api/pkg/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, response.APIResponse{Message: "ok"})
}
