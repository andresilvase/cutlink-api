package repository

import (
	"context"
	"log/slog"
	"os"

	"math/rand"

	apperrors "github.com/andresilvase/cutlink/internal/errors"
	"github.com/redis/go-redis/v9"
)

type repository struct {
	rdb *redis.Client
}

type Repository interface {
	SaveShortenedURL(ctx context.Context, fullURL string) (string, error)
	FullURL(ctx context.Context, shortenedURL string) (string, error)
}

func New(rdb *redis.Client) Repository {
	return repository{rdb}
}

var shortenderRedisKey = "shortener"

func genShortenedUrl() string {
	const caracters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJLMNOPQRSTUVWXYZ"
	const size = 8
	byts := make([]byte, size)

	for i := range size {
		randomInt := rand.Intn(len(caracters))
		byts[i] = caracters[randomInt]
	}

	return string(byts)
}

func (r repository) SaveShortenedURL(ctx context.Context, fullURL string) (string, error) {
	var shortenedURL string

	for range 100 {
		shortenedURL = genShortenedUrl()
		_, err := r.rdb.HGet(ctx, shortenderRedisKey, shortenedURL).Result()
		if err != nil {
			if err == redis.Nil {
				break
			}
			return "", err
		}
		shortenedURL = ""
	}

	if shortenedURL == "" {
		err := apperrors.LimitAttemptReached{}
		slog.Error(err.Error())
		return shortenedURL, err
	}

	_, err := r.rdb.HSet(ctx, shortenderRedisKey, shortenedURL, fullURL).Result()
	if err != nil {
		err := apperrors.Unexpected{Err: err}
		slog.Error(err.Error())
		return "", err
	}

	baseUrl := os.Getenv("SHORTENED_BASE_URL")

	return baseUrl + shortenedURL, nil
}

func (r repository) FullURL(ctx context.Context, shortenedURL string) (string, error) {
	fullURL, err := r.rdb.HGet(ctx, shortenderRedisKey, shortenedURL).Result()

	if err != nil {
		if err == redis.Nil {
			return "", apperrors.NotFound{}
		}
		unexpectedErr := apperrors.Unexpected{Err: err}
		slog.Error(unexpectedErr.Error())
		return "", unexpectedErr
	}

	return fullURL, nil
}
