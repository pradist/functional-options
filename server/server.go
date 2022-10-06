package server

import (
	"fmt"
	"time"
)

var DefaultPort = 8080

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

func New(options ...func(server *Server)) *Server {
	s := &Server{
		Port: DefaultPort,
	}
	for _, o := range options {
		o(s)
	}
	return s
}

func WithHost(host string) func(*Server) {
	return func(s *Server) {
		s.Host = host
	}
}

func WithPort(port int) func(*Server) {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConn(maxConn int) func(*Server) {
	return func(s *Server) {
		s.MaxConn = maxConn
	}
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")
	fmt.Printf("%#v", s)
	return nil
}
