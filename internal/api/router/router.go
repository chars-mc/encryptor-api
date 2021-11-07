package router

import "github.com/gorilla/mux"

type Router struct {
	Handler *mux.Router
	routes  *Routes
}

// NewRouter creates a new router and setup the routes
func NewRouter(routes *Routes) *Router {
	r := &Router{Handler: mux.NewRouter(), routes: routes}
	r.setupRoutes()
	return r
}

func (r *Router) setupRoutes() {
	for _, route := range *r.routes {
		r.Handler.HandleFunc(route.path, route.handler).Methods(route.method)
	}
}
