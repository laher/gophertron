package db

import (
	"log"

	"github.com/laher/gophertron/gophers/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionGopher = "gophers"
)

type MongoDbProvider func() (*mgo.Session, *mgo.Database, error)

type GopherDao interface {
	Spawn(gopher *model.Gopher) error
	Update(gopher *model.Gopher) error
	Die(id string) error
	GetAll() ([]model.Gopher, error)
	Get(id string) (*model.Gopher, error)
}

type GopherMongoDao struct {
	GetDb MongoDbProvider
}

func (g *GopherMongoDao) Spawn(gopher *model.Gopher) error {
	session, db, err := g.GetDb()
	if err != nil {
		log.Printf("Get Session error: %+v", err)
		return err
	}
	defer session.Close()
	cb := db.C(CollectionGopher)

	id := bson.NewObjectId()
	gopher.Id = id
	err = cb.Insert(gopher)
	return err
}

func (g *GopherMongoDao) Update(gopher *model.Gopher) error {
	session, db, err := g.GetDb()
	if err != nil {
		log.Printf("Get Session error: %+v", err)
		return err
	}
	defer session.Close()
	cb := db.C(CollectionGopher)
	err = cb.UpdateId(gopher.Id, gopher)
	return err
}

func (g *GopherMongoDao) GetAll() ([]model.Gopher, error) {
	session, db, err := g.GetDb()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	cb := db.C(CollectionGopher)
	result := []model.Gopher{}
	err = cb.Find(bson.M{}).All(&result)
	return result, err
}

func (g *GopherMongoDao) Get(gopherId string) (*model.Gopher, error) {
	session, db, err := g.GetDb()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	cb := db.C(CollectionGopher)
	result := model.Gopher{}
	err = cb.FindId(bson.ObjectIdHex(gopherId)).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (g *GopherMongoDao) Die(gopherId string) error {
	session, db, err := g.GetDb()
	if err != nil {
		return err
	}
	defer session.Close()
	cb := db.C(CollectionGopher)
	err = cb.Remove(bson.ObjectIdHex(gopherId))
	if err != nil {
		return err
	}
	return nil
}

func GetMongoDb(serverName string, dbName string) (*mgo.Session, *mgo.Database, error) {
	session, err := mgo.Dial(serverName)
	if err != nil {
		log.Printf("DB connection error for server %s: %v", serverName, err)
		return nil, nil, err
	}
	db := session.DB(dbName)
	return session, db, err
}
