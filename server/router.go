package server

import "github.com/gorilla/mux"

// Router struct
type Router struct {
	router *mux.Router
}

// New method
func (r *Router) New() *Router {
	r.router = mux.NewRouter()
	return r
}

// AddRoutes method
func (r *Router) AddRoutes() *Router {
	r.router.HandleFunc("/", HandleRoot)
	r.router.HandleFunc("/register", HandleUserRegistration)
	r.router.HandleFunc("/user", HandleUser)
	r.router.HandleFunc("/user/verify", HandleUserVerify)
	return r
}
