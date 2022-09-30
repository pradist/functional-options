package main

import (
	"github.com/pradist/functional/server"
	"log"
	"time"
)

func main() {
	svr := server.NewWithTimeoutAndMaxConn("localhost", 1234, 30*time.Second, 10)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}
