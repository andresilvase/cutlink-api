package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func SendResponse(w http.ResponseWriter, res Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(res)
	responseParseError := fmt.Errorf("error %w parsing response", err).Error()

	if err != nil {
		slog.Error(responseParseError)
		SendResponse(
			w,
			Response{Error: "Something Went Wrong! Try again later."},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)

	if _, err := w.Write(data); err != nil {
		slog.Error(responseParseError)
		return
	}
}
