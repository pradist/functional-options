package server

import (
	"fmt"
	"time"
)

type Server struct {
	Cfg Config
}

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func New(cfg Config) *Server {
	return &Server{Cfg: cfg}
}

func (s *Server) Start() error {
	fmt.Println("Starting server...")
	fmt.Printf("%#v", s)
	return nil
}
