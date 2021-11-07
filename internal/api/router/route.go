package router

import "net/http"

type Route struct {
	path    string
	method  string
	handler http.HandlerFunc
}

func NewRoute(path string, method string, handler http.HandlerFunc) *Route {
	return &Route{
		path:    path,
		method:  method,
		handler: handler,
	}
}

type Routes = []*Route
