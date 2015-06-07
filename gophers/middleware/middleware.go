package middleware

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/laher/gophertron/gophers"
	"github.com/laher/gophertron/gophers/services"
)

func MainMiddleware(authService services.AuthService, config *gophers.Config) *negroni.Negroni {
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewStatic(http.Dir("public")))
	n.Use(NewAuth(authService))
	return n
}
