package gophers

import (
	"github.com/laher/gophertron/gophers/db"
	"gopkg.in/mgo.v2"
)

type Config struct {
	DbServer    string
	DbName      string
	ServiceAddr string
}

func (config *Config) GetDb() (*mgo.Session, *mgo.Database, error) {
	return db.GetMongoDb(config.DbServer, config.DbName)
}
