package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/ipoool/laporcuranmor-api/constants"
)

// Response - to showing the response JSON
func Response(w http.ResponseWriter, httpStatus int, data interface{}, code int) {
	apiResponse := struct {
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		code,
		constants.ApiReseponse[code],
		data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(apiResponse)
}
