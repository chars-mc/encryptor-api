package server

import (
	"net/http"
	"time"

	"github.com/chars-mc/encryptor-api/internal/api/router"
)

type Server struct {
	Port   string
	router *router.Router
	http   *http.Server
}

func NewServer(port string, router *router.Router) *Server {
	return &Server{
		Port:   port,
		router: router,
		http: &http.Server{
			Addr:           port,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        router.Handler,
		},
	}
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}
