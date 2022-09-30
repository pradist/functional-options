package server

import (
	"fmt"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

func New(host string, port int) *Server {
	return &Server{host, port, time.Minute, 100}
}

func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{host: host, port: port, timeout: timeout}
}

func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{host, port, timeout, maxConn}
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")
	fmt.Printf("%#v", s)
	return nil
}
