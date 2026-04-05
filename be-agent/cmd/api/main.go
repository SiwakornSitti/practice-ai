package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SiwakornSitti/practice-ai/be-agent/internal/user/delivery"
	"github.com/SiwakornSitti/practice-ai/be-agent/internal/user/repository"
	"github.com/SiwakornSitti/practice-ai/be-agent/internal/user/usecase"
)

func main() {
	// Secret key for JWT. In production, load from environment variable
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-key-for-development-only"
	}

	// Initialize Repository
	userRepo := repository.NewMemoryUserRepository()

	// Initialize Usecase
	userUsecase := usecase.NewUserUsecase(userRepo, jwtSecret)

	// Initialize standard library Mux
	mux := http.NewServeMux()

	// Register Handlers
	delivery.NewUserHandler(mux, userUsecase, jwtSecret)

	// Add a simple health check route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "be-agent API is running. Clean Architecture implementation.")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
