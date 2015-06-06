package main

import (
	"log"

	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/wiring"
)

func main() {
	log.Print("Starting gophertron")
	//TODO: flags for config
	config := &gophers.Config{
		DbName:      "gophertron",
		DbServer:    "localhost",
		ServiceAddr: ":8001",
	}
	wiring.Wiring(config)
}
