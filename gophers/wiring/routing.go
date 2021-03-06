package wiring

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/model"
	"github.com/laher/gophertron/gophers/webapi"
)

func routing(gopherApi webapi.GopherApi, config *gophers.Config) http.Handler {
	wsContainer := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Path("/gophers").
		Doc("Bother gophers").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Route(ws.GET("/").To(gopherApi.GetAll).
		// docs
		Doc("get all gophers").
		Operation("allGopher").
		Writes([]model.Gopher{})) // on the response
	ws.Route(ws.GET("/{gopherId}").To(gopherApi.GetGopher).
		// docs
		Doc("get a gopher").
		Operation("findGopher").
		Param(ws.PathParameter("gopherId", "identifier of the gopher").DataType("string")).
		Writes(model.Gopher{})) // on the response

	ws.Route(ws.PUT("/{gopherId}").To(gopherApi.Put).
		// docs
		Doc("update a gopher").
		Operation("updateGopher").
		Param(ws.PathParameter("gopherId", "identifier of the gopher").DataType("string")).
		Reads(model.Gopher{}).  // from the request
		Writes(model.Gopher{})) // on the response

	ws.Route(ws.POST("/{gopherId}/zap").To(gopherApi.Zap).
		// docs
		Doc("zap a gopher. Removes its skillz").
		Operation("zapGopher").
		Param(ws.PathParameter("gopherId", "identifier of the gopher").DataType("string")).
		Writes(model.Gopher{})) // on the response

	ws.Route(ws.POST("/{gopherId}/kapow").To(gopherApi.Kapow).
		// docs
		Doc("kapow a gopher. Removes its skillz. If it has no skillz, returns error").
		Operation("zapGopher").
		Param(ws.PathParameter("gopherId", "identifier of the gopher").DataType("string")).
		Writes(model.Gopher{})) // on the response

	ws.Route(ws.POST("").To(gopherApi.Post).
		// docs
		Doc("create a gopher").
		Operation("createGopher").
		Reads(model.Gopher{}).  // from the request
		Writes(model.Gopher{})) // on the response

	ws.Route(ws.DELETE("/{gopherId}").To(gopherApi.Delete).
		// docs
		Doc("delete a gopher").
		Operation("removeGopher").
		Param(ws.PathParameter("gopherId", "identifier of the gopher").DataType("string")))

	wsContainer.Add(ws)
	swConfig := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "http://localhost" + config.ServiceAddr,
		ApiPath:        "/apidocs.json",
		// Optionally, specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "../swagger-ui/dist"} //download to here ...
	swagger.RegisterSwaggerService(swConfig, wsContainer)

	return wsContainer
}
