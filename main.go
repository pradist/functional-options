package main

import (
	"github.com/pradist/functional/server"
	"log"
	"time"
)

func main() {
	svr := server.New(server.Config{Host: "localhost", Port: 1234, Timeout: 30 * time.Second, MaxConn: 10})
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
