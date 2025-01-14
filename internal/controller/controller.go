package controller

import (
	"context"
	"log/slog"
	"net/url"
	"strings"

	apperrors "github.com/andresilvase/cutlink/internal/errors"
	"github.com/andresilvase/cutlink/internal/repository"
	"github.com/redis/go-redis/v9"
)

var store = repository.New(
	redis.NewClient(&redis.Options{Addr: "redis:6379"}),
)

func isValideURL(fullURL string) bool {
	tempURL := fullURL
	if !strings.Contains(tempURL, "://") {
		tempURL = "http://" + tempURL
	}

	parsedUrl, err := url.Parse(tempURL)
	return err == nil && parsedUrl.IsAbs()
}

func ShortenURL(fullURL string) (string, error) {
	slog.Info(fullURL)
	if !isValideURL(fullURL) {
		return "", apperrors.InvalideURL{}
	}

	ctx := context.Background()

	return store.SaveShortenedURL(ctx, fullURL)
}

func FullURL(shortenedURL string) (string, error) {
	ctx := context.Background()

	return store.FullURL(ctx, shortenedURL)
}
