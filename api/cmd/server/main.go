package main

import (
	"log"
	"net/http"

	"api/internal/config"
	"api/internal/router"
	"api/internal/services"
)

func main() {
	cfg := config.Load()
	authService := services.NewAuthService(cfg.JWTSecret)
	r := router.New(authService)

	addr := ":" + cfg.Port
	log.Printf("API server running on %s", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
