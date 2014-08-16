package gophers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"time"
)

// API for handling gopher-related requests
type GopherApi struct {
	Dao GopherDao
}

func loadGopherFromJson(r io.Reader) (*Gopher, error) {
	var gopher Gopher
	jsonBlob, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("Read body error: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(jsonBlob, &gopher)
	if err != nil {
		log.Printf("Unmarshal error: %+v", err)
		log.Printf("Blob: %+v", string(jsonBlob))
		return nil, err
	}
	return &gopher, err
}

func (g *GopherApi) Post(w http.ResponseWriter, r *http.Request) {
	gopher, err := loadGopherFromJson(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//New: set Born to now
	gopher.Born = time.Now()
	gopher.Mutated = time.Now()
	err = g.Dao.Spawn(gopher)
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	marshalJson(w, gopher) //gopher will receive updates ...
}

func (g *GopherApi) Put(w http.ResponseWriter, r *http.Request) {
	gopher, err := loadGopherFromJson(r.Body)
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gopher.Mutated = time.Now()
	err = g.Dao.Update(gopher)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Printf("Gopher not found: %+v", gopher)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	marshalJson(w, gopher) //gopher will receive updates ...

}
func (g *GopherApi) Get(w http.ResponseWriter, params map[string]string) {
	gopher, err := g.Dao.Get(params["gopherId"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("About to return gopher %+v", gopher)
	marshalJson(w, gopher)
}

func (g *GopherApi) Zap(w http.ResponseWriter, params map[string]string) {
	gopher, err := g.Dao.Get(params["gopherId"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	gopher.Zap()
	err = g.Dao.Update(gopher)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Printf("Gopher not found: %+v", gopher)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("About to return gopher %+v", gopher)
	marshalJson(w, gopher)
}

func (g *GopherApi) Kapow(w http.ResponseWriter, params map[string]string) {
	gopher, err := g.Dao.Get(params["gopherId"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	gopher.Kapow()
	err = g.Dao.Update(gopher)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Printf("Gopher not found: %+v", gopher)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	log.Printf("About to return gopher %+v", gopher)
	marshalJson(w, gopher)
}

func (g *GopherApi) GetAll(w http.ResponseWriter) {
	gophers, err := g.Dao.GetAll()
	if err != nil {
		log.Printf("Server error: %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("About to return all gophers %+v", gophers)
	marshalJson(w, gophers)

}

func (g *GopherApi) Delete(w http.ResponseWriter, params map[string]string) {
	err := g.Dao.Die(params["gopherId"])
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if err != nil {
			log.Printf("Server error: %+v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, "Deleted OK", http.StatusNoContent)
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
