package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run starts the http server
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func NewServer() *Server {
	//log.Info("server started")
	r := mux.NewRouter()

	//TODO: add configurations
	address := "8080"
	server := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return &Server{httpServer: server}
}
