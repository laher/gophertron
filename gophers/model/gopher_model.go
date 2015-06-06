package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Gopher struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string
	Skillz  []string
	Born    time.Time
	Mutated time.Time
}

//Zap removes a gopher's skillz
func (gopher *Gopher) Zap() {
	gopher.Skillz = []string{}
	gopher.Mutated = time.Now()
}

//Pow removes a gopher's skills. However, if a gopher already has no skills, Kapow returns an error
// If gopher already has no skills, then return an error
//NOTE: this is used inside gophertron as an example of incorrect pointer usage
func (gopher Gopher) Kapow() error {
	if gopher.Skillz == nil || len(gopher.Skillz) == 0 {
		return fmt.Errorf("This gopher can't take any more Kapows, capn")
	}
	gopher.Skillz = []string{}
	gopher.Mutated = time.Now()
	return nil
}
