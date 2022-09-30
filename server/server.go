package server

import "fmt"

type Server struct {
	host string
	port int
}

func New(host string, port int) *Server {
	return &Server{host, port}
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")
	return nil
}
