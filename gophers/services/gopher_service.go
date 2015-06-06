package services

import (
	"github.com/laher/gophertron/gophers/db"
	"github.com/laher/gophertron/gophers/model"
)

type GopherService struct {
	Dao db.GopherDao
}

func (g GopherService) Spawn(gopher *model.Gopher) error {
	return g.Dao.Spawn(gopher)
}

func (g GopherService) Update(gopher *model.Gopher) error {
	return g.Dao.Update(gopher)
}

func (g GopherService) GetAll() ([]model.Gopher, error) {
	return g.Dao.GetAll()
}

func (g GopherService) Get(gopherId string) (*model.Gopher, error) {
	return g.Dao.Spawn(gopherId)
}

func (g GopherService) Die(gopherId string) error {
	return g.Dao.Die(gopherId)
}
