package api

import (
	"log"
	"net/http"
	"os"
	"strings"

	routes "github.com/andresilvase/cutlink/cmd/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

func handler() http.Handler {
	router := chi.NewMux()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

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
	router.Use(corsHandler.Handler)

	router.Route("/", func(r chi.Router) {
		r.Get("/{shortenedUrl:[a-zA-Z0-9]+}", routes.FullURL)
		r.Get("/health", routes.HealthCheck)
		r.Post("/cut", routes.ShortenLink)
	})

	return router
}
