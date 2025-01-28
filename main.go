package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"fiter/models"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/", fs)

	// Band-related endpoints
	 http.HandleFunc("/bands", models.BandsHandler)// this  handler displays the data of all artsists
	http.HandleFunc("/filter", models.FilterBandsHandler)// displays data of fltered artists only	

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Start server with graceful shutdown
	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Server listening on http://localhost:8080/")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
