package webapi

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/laher/gophertron/gophers/model"
	"gopkg.in/mgo.v2/bson"
)

//newRecorder creates an httptest.NewRecorder, wraps this and an http.Request inside restful.Request and restful.Response respectively.
func newRecorder(r *http.Request) (*restful.Request, *restful.Response, *httptest.ResponseRecorder) {
	restful.DefaultResponseContentType(restful.MIME_JSON)
	w := httptest.NewRecorder()
	r.Header.Add("Accept", "application/json, text/plain, text/html")
	request := restful.NewRequest(r)
	response := restful.NewResponse(w)
	return request, response, w
}

//gopherDummyDAO doesnt send errors.
type gopherDummyDAO struct {
}

func (g *gopherDummyDAO) Spawn(gopher *model.Gopher) error {
	gopher.Id = bson.NewObjectId()
	return nil
}
func (g *gopherDummyDAO) Update(gopher *model.Gopher) error {
	return nil
}
func (g *gopherDummyDAO) GetAll() ([]model.Gopher, error) {
	return []model.Gopher{}, nil
}

func (g *gopherDummyDAO) Get(gopherId string) (*model.Gopher, error) {
	return &model.Gopher{Id: bson.ObjectIdHex(gopherId), Name: "roger", Skillz: []string{"Fencing", "Badger"}}, nil
}

func (g *gopherDummyDAO) Die(gopherId string) error {
	return nil
}

func TestApiPost(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	rdr := strings.NewReader(`{ "name": "diane" }`)
	r, err := http.NewRequest("POST", "/gophers", rdr)
	r.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	gopherApi.Post(req, resp)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiPostBadRequest(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	rdr := strings.NewReader(`{ "name": "diane }`)
	r, err := http.NewRequest("POST", "/gophers", rdr)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	gopherApi.Post(req, resp)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiGet(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	gopherId := bson.NewObjectId().Hex()
	t.Logf("Gopher id %s", gopherId)
	r, err := http.NewRequest("GET", "/gophers/"+gopherId, nil)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	req.PathParameters()["gopherId"] = gopherId
	gopherApi.GetGopher(req, resp)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiGetAll(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	r, err := http.NewRequest("GET", "/gophers/", nil)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	gopherApi.GetAll(req, resp)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiDelete(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	gopherId := bson.NewObjectId().Hex()
	r, err := http.NewRequest("DELETE", "/gophers/"+gopherId, nil)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	req.PathParameters()["gopherId"] = gopherId
	gopherApi.Delete(req, resp)
	if w.Code != http.StatusNoContent {
		t.Fatalf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}

func TestApiKapow(t *testing.T) {
	gopherApi := &GopherApi{Dao: &gopherDummyDAO{}}
	gopherId := bson.NewObjectId().Hex()
	r, err := http.NewRequest("POST", "/gophers/"+gopherId+"/kapow", nil)
	if err != nil {
		t.Fatalf("request error: %v", err)
	}
	req, resp, w := newRecorder(r)
	req.PathParameters()["gopherId"] = gopherId
	gopherApi.Kapow(req, resp)
	if w.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", w.Code)
	}
	t.Logf("Status: %d, Response data: %v", w.Code, w.Body)
}
