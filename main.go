package main

import (
	"log"

	"github.com/pradist/functional-options/server"
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
