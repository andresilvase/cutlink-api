package api

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/andresilvase/cutlink/cmd/api/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func handler() http.Handler {
	handler := chi.NewMux()

	handler.Use(middleware.Recoverer)
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)

	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")

	if allowedOriginsEnv == "" {
		log.Fatal("ALLOWED_ORIGINS is not set in the environment")
	}

	allowedOrigins := strings.Split(allowedOriginsEnv, ",")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler.Use(corsHandler.Handler)

	handler.Route("/", func(r chi.Router) {
		r.Get("/{shortenedUrl:[a-zA-Z0-9]+}", routes.FullURL)
		r.Get("/health", routes.HealthCheck)
		r.Post("/cut", routes.ShortenLink)
	})

	return handler
}
