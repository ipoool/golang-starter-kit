package interfaces

import "net/http"

// IUserHandler - UserHandler interface
type IUserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}
