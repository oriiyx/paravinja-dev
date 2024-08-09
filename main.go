package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	router2 "github.com/oriiyx/paravinja-dev/internal/router"
)

func main() {
	initialize()
	ctx := context.Background()

	router := chi.NewRouter()
	routerController := router2.Controller{
		Router: router,
	}

	routerController.RegisterUses()
	routerController.RegisterRoutes()

	server := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	log.Println("Starting server on: http://localhost:80")

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("Server Exited Properly")
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
