package main

import (
	"github.com/go-martini/martini"
	"github.com/laher/gophertron/gophers"
	"labix.org/v2/mgo"
	"log"
	"net/http"
)

type Config struct {
	DbServer    string
	DbName      string
	ServiceAddr string
}

func (config *Config) GetDb() (*mgo.Session, *mgo.Database, error) {
	return gophers.GetMongoDb(config.DbServer, config.DbName)
}

type parami func(w http.ResponseWriter, params map[string]string)
type paramo func(w http.ResponseWriter, params martini.Params)
func pw(i parami) paramo {
	return func(w http.ResponseWriter, params martini.Params) {
		i(w, map[string]string(params))
	}
}

func main() {
	log.Print("Starting gophertron")
	config := &Config{
		DbName:      "gophertron",
		DbServer:    "localhost",
		ServiceAddr: ":8001",
	}
	router := martini.Classic()
	dao := &gophers.GopherMongoDao{GetDb: config.GetDb}
	gopherApi := &gophers.GopherApi{dao}
	router.Get("/gophers", gopherApi.GetAll)
	//new
	router.Post("/gophers", gopherApi.Post)
	router.Post("/gophers/:gopherId/zap", pw(gopherApi.Zap))
	router.Post("/gophers/:gopherId/kapow", pw(gopherApi.Kapow))
	//existing
	router.Put("/gophers/", gopherApi.Put)
	router.Get("/gophers/:gopherId", pw(gopherApi.Get))
	router.Delete("/gophers/:gopherId", pw(gopherApi.Delete))
	http.Handle("/", router)

	log.Printf("Gophertron to listen on %s", config.ServiceAddr)
	err := http.ListenAndServe(config.ServiceAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
