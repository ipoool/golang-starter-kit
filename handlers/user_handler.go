package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ipoool/golang-starter-kit/constants"
	"github.com/ipoool/golang-starter-kit/helpers"
	"github.com/ipoool/golang-starter-kit/interfaces"
)

// UserHandler - User Handler struct
type UserHandler struct {
	userRepository interfaces.IUserRepository
}

// GetUser - For handling get  user
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := helpers.StringToInt(vars["user_id"])

	// get user by id from user repository
	user, err := h.userRepository.GetUserByID(userID)
	if err != nil {
		helpers.Response(w, http.StatusOK, nil, constants.ApiFailedMessage)
		return
	}

	helpers.Response(w, http.StatusOK, user, constants.ApiSuccessMessage)
	return
}
