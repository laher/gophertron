package webapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	"github.com/laher/gophertron/gophers/model"
	"github.com/laher/gophertron/gophers/services"
	"gopkg.in/mgo.v2"
)

// API for handling gopher-related requests
type GopherApi struct {
	GopherService services.GopherService
}

func (g GopherApi) Post(request *restful.Request, response *restful.Response) {
	gopher := new(model.Gopher)
	err := request.ReadEntity(gopher)
	if err != nil {
		http.Error(response.ResponseWriter, err.Error(), http.StatusBadRequest)
	} else {
		//New: set Born to now
		gopher.Born = time.Now()
		gopher.Mutated = time.Now()
		err = g.GopherService.Spawn(gopher)
		if err != nil {
			log.Printf("API returning server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
		} else {
			response.AddHeader("Content-Type", "application/json")
			response.WriteEntity(gopher)
		}
	}
}

func (g GopherApi) Put(request *restful.Request, response *restful.Response) {
	gopher := new(model.Gopher)
	err := request.ReadEntity(gopher)
	if err != nil {
		http.Error(response.ResponseWriter, err.Error(), http.StatusBadRequest)
	} else {
		gopher.Mutated = time.Now()
		err = g.GopherService.Update(gopher)
		if err != nil {
			if err == mgo.ErrNotFound {
				log.Printf("Gopher not found: %+v", gopher)
				http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			} else if err != nil {
				log.Printf("API returning server error: %+v", err)
				http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			}

		} else {
			response.AddHeader("Content-Type", "application/json")
			response.WriteEntity(gopher)
		}
	}
}

func (g GopherApi) GetGopher(request *restful.Request, response *restful.Response) {
	gopherId := request.PathParameter("gopherId")
	gopher, err := g.GopherService.Get(gopherId)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Printf("About to return gopher %+v", gopher)
		response.WriteEntity(gopher)
	}
}

func (g GopherApi) Zap(request *restful.Request, response *restful.Response) {
	gopherId := request.PathParameter("gopherId")
	gopher, err := g.GopherService.Get(gopherId)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	gopher.Zap()
	err = g.GopherService.Update(gopher)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Printf("Gopher not found: %+v", gopher)
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("About to return gopher %+v", gopher)
	response.AddHeader("Content-Type", "application/json")
	response.WriteEntity(gopher)
}

func (g GopherApi) Kapow(request *restful.Request, response *restful.Response) {
	gopherId := request.PathParameter("gopherId")
	gopher, err := g.GopherService.Get(gopherId)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	err = gopher.Kapow()
	if err != nil {
		log.Printf("Kapow error: %+v", err)
		http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	err = g.GopherService.Update(gopher)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Printf("Gopher not found: %+v", gopher)
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("About to return gopher %+v", gopher)
	response.AddHeader("Content-Type", "application/json")
	response.WriteEntity(gopher)
}

func (g GopherApi) GetAll(request *restful.Request, response *restful.Response) {
	gophers, err := g.GopherService.GetAll()
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("About to return all gophers %+v", gophers)
	response.AddHeader("Content-Type", "application/json")
	response.WriteEntity(gophers)
}

func (g GopherApi) Delete(request *restful.Request, response *restful.Response) {
	gopherId := request.PathParameter("gopherId")
	err := g.GopherService.Die(gopherId)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(response.ResponseWriter, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(response.ResponseWriter, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Error(response.ResponseWriter, "Deleted OK", http.StatusNoContent)
}

func marshalJson(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(result)
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(b))
}
