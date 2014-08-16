package gophers

import(
	"labix.org/v2/mgo/bson"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

//gopherDummyDAO doesnt send errors.
type gopherDummyDAO struct {
}
func (g *gopherDummyDAO) Spawn(gopher *Gopher) error {
	gopher.Id = bson.NewObjectId()
	return nil
}
func (g *gopherDummyDAO) Update(gopher *Gopher) error {
	return nil
}
func (g *gopherDummyDAO) GetAll() ([]Gopher, error) {
	return []Gopher{}, nil
}
func (g *gopherDummyDAO) Get(gopherId string) (*Gopher, error) {
	return &Gopher{Id:bson.ObjectIdHex(gopherId), Name: "roger", Skillz: []string{"Fencing", "Badger"}}, nil
}
func (g *gopherDummyDAO) Die(gopherId string) error {
	return nil
}



func TestApiPost(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	r := strings.NewReader(`{ "name": "diane" }`)
	req, err := http.NewRequest("POST", "/gophers", r)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	gopherApi.Post(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiPostBadRequest(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	r := strings.NewReader(`{ "name": "diane }`)
	req, err := http.NewRequest("POST", "/gophers", r)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}

	gopherApi.Post(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiGet(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	params := make(map[string]string)
	params["gopherId"] = bson.NewObjectId().Hex()
	gopherApi.Get(w, params)

	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiGetAll(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	gopherApi.GetAll(w)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiDelete(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	params := make(map[string]string)
	params["gopherId"] = bson.NewObjectId().Hex()
	gopherApi.Delete(w, params)

	if w.Code != http.StatusNoContent {
		t.Fatalf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiKapow(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	w := httptest.NewRecorder()
	params := make(map[string]string)
	params["gopherId"] = bson.NewObjectId().Hex()
	gopherApi.Kapow(w, params)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}
