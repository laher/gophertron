package gophers

import(
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
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
	w := httptest.NewRecorder()
	r := strings.NewReader(`{ "name": "diane" }`)
	req, err := http.NewRequest("POST", "/gophers", r)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	gopherApi.Post(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)

}