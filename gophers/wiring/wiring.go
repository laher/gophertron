package wiring

import (
	"net/http"

	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/db"
	"github.com/laher/gophertron/gophers/services"
	"github.com/laher/gophertron/gophers/webapi"
)

func Wiring(config *gophers.Config) http.Handler {

	dao := db.GopherMongoDao{GetDb: config.GetDb}
	service := services.GopherService{Dao: dao}
	gopherApi := webapi.GopherApi{GopherService: service}

	return routing(gopherApi, config)
}
