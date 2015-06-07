package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/laher/gophertron/gophers/services"
)

func NewAuth(authService services.AuthService) negroni.Handler {
	m := BasicAuthMiddleware{AuthService: authService, Realm: "gopher-cave"}
	return m
}

type BasicAuthMiddleware struct {
	AuthService services.AuthService
	//UpHashed    map[string]string
	Realm string
}

//WARNING this is not fully tested. We don't use this in production.
//(we use oauth)
func (bam BasicAuthMiddleware) ServeHTTP(resp http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authorizationArray := req.Header["Authorization"]

	if len(authorizationArray) > 0 {
		authorization := strings.TrimSpace(authorizationArray[0])
		credentials := strings.Split(authorization, " ")

		if len(credentials) != 2 || credentials[0] != "Basic" {
			logrus.Warnf("Error getting credentials (not basic auth)")
			unauthorized(resp, bam.Realm)
			return
		}

		authstr, err := base64.StdEncoding.DecodeString(credentials[1])
		if err != nil {
			logrus.Warnf("Error decoding credentials %v", err)
			unauthorized(resp, bam.Realm)
			return
		}

		userpass := strings.Split(string(authstr), ":")
		if len(userpass) != 2 {
			logrus.Warnf("Error getting u/p %v", err)
			unauthorized(resp, bam.Realm)
			return
		}
		isAuth := bam.AuthService.Auth(userpass[0], userpass[1])
		if !isAuth {
			logrus.Warnf("Non-matching u/p %v", err)
			unauthorized(resp, bam.Realm)
			return
		}
		next(resp, req)
	} else {
		logrus.Warnf("No user info")
		unauthorized(resp, bam.Realm)
	}
}

func unauthorized(w http.ResponseWriter, realm string) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf("Basic realm=\"%s\"", realm))
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
