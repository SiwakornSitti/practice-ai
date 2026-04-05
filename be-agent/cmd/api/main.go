package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Run server in a goroutine so that it doesn't block
	go func() {
		log.Printf("Server starting on :%s...", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
