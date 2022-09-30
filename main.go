package main

import (
	"github.com/pradist/functional/server"
	"log"
)

func main() {
	svr := server.New("localhost", 8080)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
