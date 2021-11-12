package router

import (
	"net/http"

	"github.com/chars-mc/encryptor-api/internal/api/middlewares"
	"github.com/gorilla/mux"
)

type Router struct {
	Handler *mux.Router
	routes  *Routes
}

// NewRouter creates a new router and setup the routes
func NewRouter(routes *Routes) *Router {
	r := &Router{Handler: mux.NewRouter(), routes: routes}
	r.setupRoutes()
	r.setupMiddlewares()
	return r
}

func (r *Router) setupRoutes() {
	for _, route := range *r.routes {
		r.Handler.HandleFunc(route.path, route.handler).Methods(route.method, http.MethodOptions)
	}
}

func (r *Router) setupMiddlewares() {
	r.Handler.Use(middlewares.Logger)
	r.Handler.Use(middlewares.Cors)
}
