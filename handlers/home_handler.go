package handlers

import (
	"net/http"

	"github.com/ipoool/laporcuranmor-api/constants"
	"github.com/ipoool/laporcuranmor-api/helpers"
)

// HomeHandler - Handling for home
type HomeHandler struct{}

// HeartBeat - For checking status service
func (h *HomeHandler) HeartBeat(w http.ResponseWriter, r *http.Request) {

	helpers.Response(w, http.StatusOK, nil, constants.ApiSuccessMessage)
}
