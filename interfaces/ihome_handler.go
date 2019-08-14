package interfaces

import "net/http"

// IHomeHandler - Interface for HomeHandler
type IHomeHandler interface {
	HeartBeat(w http.ResponseWriter, r *http.Request)
}
