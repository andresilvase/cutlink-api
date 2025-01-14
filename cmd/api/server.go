package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Run() {
	handler := handler()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		errorMsg := fmt.Errorf("there was an error initializing the server %w", err)
		slog.Error(errorMsg.Error())
		panic(errorMsg)
	}
}
