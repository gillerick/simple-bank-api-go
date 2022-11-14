package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"net/http"
	"simple-bank-account/configs"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Run starts the http server
func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

// NewServer sets up a new server using the specified configurations
func NewServer(config configs.App, customerHandler, accountHandler, cardHandler http.Handler) *Server {
	log.Infof("server listening on address: %s, port: %s", config.Host, config.Port)
	r := mux.NewRouter()

	//ToDo: Setup routes

	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	server := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return &Server{httpServer: server}
}
