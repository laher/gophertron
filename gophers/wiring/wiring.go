package wiring

import (
	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/db"
	"github.com/laher/gophertron/gophers/services"
	"github.com/laher/gophertron/gophers/webapi"
)

func Wiring(config *gophers.Config) {

	dao := db.GopherMongoDao{GetDb: config.GetDb}
	service := services.GopherService{Dao: dao}
	gopherApi := webapi.GopherApi{GopherService: service}

	routing(gopherApi, config)
}
