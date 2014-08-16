package gophers

import (
	"time"
	"fmt"
	"labix.org/v2/mgo/bson"
)
/*
type Id string

func (id *Id) Hex() string {
	dest := make([]byte, len(*id)*2)
	hex.Encode(dest, []byte(*id))
	return string(dest)
}
func (id *Id) String() string {
	return id.Hex()
}
func (id *Id) MarshalJSON() ([]byte, error) {
	hex := id.Hex()
	quotedDestString := fmt.Sprintf(`"%s"`, hex)
	return []byte(quotedDestString), nil
}
func (id *Id) UnmarshalJSON(input []byte) error {
	if string(input) == "null" {
		return nil
	}
	if len(input) < 2  {
		return fmt.Errorf("string too short")
	}
	in := input[1:len(input)-1]
	dest := make([]byte, len(in)/2)
	_, err := hex.Decode(dest, in)
	if err != nil {
		return err
	}
	*id = Id(string(dest))
	fmt.Println("printed", *id)
	return nil
}
*/

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
