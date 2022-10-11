package main

import (
	"log"
	"time"

	"github.com/pradist/functional-options/server"
)

func main() {
	svr := server.New(server.Config{
		Host:    "localhost",
		Port:    1234,
		Timeout: 30 * time.Second,
	})

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
