package gophers

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

//gopherErrorDAO returns errors for any call
type gopherErrorDAO struct {
}

func (g *gopherErrorDAO) Spawn(gopher *Gopher) error {
	return fmt.Errorf("Invalid input")
}
func (g *gopherErrorDAO) Update(gopher *Gopher) error {
	return fmt.Errorf("invalid gopher")
}
func (g *gopherErrorDAO) GetAll() ([]Gopher, error) {
	return nil, fmt.Errorf("The gophers are revolting")
}
func (g *gopherErrorDAO) Get(gopherId string) (*Gopher, error) {
	return nil, fmt.Errorf("No gopher found")
}

func (g *gopherErrorDAO) Die(gopherId string) error {
	return fmt.Errorf("Gopher already dead")
}

func TestApiPostServerError(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherErrorDAO{}}
	rdr := strings.NewReader(`{ "name": "diane" }`)
	r, err := http.NewRequest("POST", "/gophers", rdr)
	r.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	gopherApi.Post(req, resp)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)

}
