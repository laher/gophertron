package gophers

import (
	"os"

	"gopkg.in/mgo.v2"

	"net/http"
)

func IsNotFoundError(err error) (bool, string) {
	if err == mgo.ErrNotFound {
		return true, "Entity not found"
	}
	if os.IsNotExist(err) {
		return true, "file not found"
	}
	_, ok := err.(ErrorNotFound)
	if ok {
		return true, err.Error()
	}

	return false, ""
}

func IsBadRequest(err error) (bool, string) {
	if mgo.IsDup(err) {
		return true, "Key already exists"
	}
	_, ok := err.(ErrorBadRequest)
	if ok {
		return true, err.Error()
	}
	return false, ""
}

// Checks for some known errors
func GetErrorHttpStatus(err error) (string, int) {

	if err == nil {
		return "OK", http.StatusOK
	}

	_, ok := err.(ErrorForbidden)
	if ok {
		return err.Error(), http.StatusForbidden
	}

	_, ok = err.(ErrorUnauthorized)
	if ok {
		return err.Error(), http.StatusUnauthorized
	}

	ok, msg := IsBadRequest(err)
	if ok {
		return msg, http.StatusBadRequest
	}

	ok, msg = IsNotFoundError(err)
	if ok {
		return err.Error(), http.StatusNotFound
	}
	//
	return err.Error(), http.StatusInternalServerError
}

type ErrorInternalServerError struct {
	Err    error  //default
	AltMsg string //alternate message
}

func (e ErrorInternalServerError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.AltMsg
}

type ErrorUnauthorized struct {
	Err    error  //default
	AltMsg string //alternate message
}

func (e ErrorUnauthorized) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.AltMsg
}

type ErrorForbidden struct {
	Err    error  //default
	AltMsg string //alternate message
}

func (e ErrorForbidden) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.AltMsg
}

type ErrorNotFound struct {
	Err    error  //default
	AltMsg string //alternate message
}

func (e ErrorNotFound) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.AltMsg
}

type ErrorBadRequest struct {
	Err    error  //default
	AltMsg string //alternate message
}

func (e ErrorBadRequest) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.AltMsg
}
