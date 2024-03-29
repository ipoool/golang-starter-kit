package routes

import (
	"github.com/gorilla/mux"
	"github.com/ipoool/golang-starter-kit/handlers"
	"github.com/ipoool/golang-starter-kit/interfaces"
	"github.com/urfave/negroni"
)

// Route - Routing Services
type Route struct {
	Redis       interfaces.IRedis
	DB          interfaces.ISQL
	middleware  Middleware
	homeHandler interfaces.IHomeHandler
	userHandler interfaces.IUserHandler
}

// Init - Initialize the route
func (s *Route) Init() {
	s.homeHandler = &handlers.HomeHandler{}
	s.userHandler = &handlers.UserHandler{}
}

// Handle - Handling route
func (s *Route) Handle() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	var versionAPI = router.PathPrefix("/v1").Subrouter()
	var externalAPI = versionAPI.PathPrefix("/ex").Subrouter()

	versionAPI.HandleFunc("/heartbeat", s.homeHandler.HeartBeat).Methods("GET")
	externalAPI.HandleFunc("/user/{user_id}", s.userHandler.GetUser).Methods("GET")

	allRoute := mux.NewRouter()
	externalAPI.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		allRoute.Path(t).Handler(negroni.New(
			negroni.HandlerFunc(s.middleware.Auth),
			negroni.Wrap(externalAPI),
		))
		return nil
	})
	allRoute.PathPrefix("/").Handler(router)

	return allRoute
}
