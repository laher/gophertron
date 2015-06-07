package main

import (
	"log"
	"net/http"

	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/middleware"
	"github.com/laher/gophertron/gophers/services"
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
	wsContainer := wiring.Wiring(config)
	authService := services.DummyAuthService{}
	mw := middleware.MainMiddleware(authService, config)
	mw.UseHandler(wsContainer)
	http.Handle("/", mw)
	log.Printf("Gophertron to listen on %s", config.ServiceAddr)
	err := http.ListenAndServe(config.ServiceAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
