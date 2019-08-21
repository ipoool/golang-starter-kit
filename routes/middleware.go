package routes

import (
	"net/http"

	"github.com/ipoool/laporcuranmor-api/constants"
	"github.com/ipoool/laporcuranmor-api/helpers"
)

// Middleware - Middleware api
type Middleware struct{}

// Auth - Middleware authentication
func (m *Middleware) Auth(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var token = r.Header.Get("Authorization")

	jwt := new(helpers.JWT)
	if err := jwt.CheckToken(token); err != nil {
		helpers.Response(w, http.StatusUnauthorized, nil, constants.ApiAuthorizationFailed)
		return
	}
	next(w, r)
}
