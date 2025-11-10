package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port string
}

func New(port string) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", HealthHandler)
	mux.HandleFunc("/validate", ValidateHandler)

	addr := fmt.Sprintf(":%s", s.port)
	log.Printf("Server starting on %s", addr)

	return http.ListenAndServe(addr, mux)
}
