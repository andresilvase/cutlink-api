package routes

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/andresilvase/cutlink/cmd/api/utils"
	"github.com/andresilvase/cutlink/internal/controller"
	apperrors "github.com/andresilvase/cutlink/internal/errors"
	"github.com/andresilvase/cutlink/internal/models"
	"github.com/go-chi/chi/v5"
)

func getFullURL(r *http.Request) string {
	scheme := "http"

	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s\n", scheme, r.Host, r.RequestURI)
}

func FullURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	shortenedUrlParameter := chi.URLParam(r, "shortenedUrl")

	pathRequest := getFullURL(r)
	isAPIEndpoint := strings.Contains(pathRequest, "api")

	fmt.Println(shortenedUrlParameter)
	fmt.Println(pathRequest)

	fullURL, err := controller.FullURL(shortenedUrlParameter)
	baseUrl := os.Getenv("SHORTENED_BASE_URL")

	if err != nil {
		statusCode := http.StatusInternalServerError
		msgError := err.Error()
		slog.Error(msgError)

		if errors.As(err, &apperrors.NotFound{}) {
			statusCode = http.StatusNotFound
		}

		if isAPIEndpoint {
			utils.SendResponse(
				w,
				utils.Response{Error: msgError},
				statusCode,
			)

		} else {
			http.Redirect(
				w,
				r,
				fmt.Sprintf("%s%s", baseUrl, "not-found/"),
				http.StatusMovedPermanently,
			)
		}

		return
	}

	response := models.FullURL{
		URL: fullURL,
	}

	if isAPIEndpoint {
		utils.SendResponse(
			w,
			utils.Response{
				Data: response,
			},
			http.StatusOK,
		)
	} else {
		http.Redirect(
			w,
			r,
			response.URL,
			http.StatusFound,
		)
	}

	return
}
