package main

import (
	"github.com/pradist/functional/server"
	"log"
)

func main() {
	svr := server.New(
		server.WithHost("localhost"),
		server.WithTimeout(10),
	)

	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
