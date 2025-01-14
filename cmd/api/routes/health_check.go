package routes

import (
	"net/http"

	"github.com/andresilvase/cutlink/cmd/api/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	utils.SendResponse(
		w,
		utils.Response{
			Data: "Service Healthy",
		},
		http.StatusOK,
	)
}
