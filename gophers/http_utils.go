package gophers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Sirupsen/logrus"

	"io/ioutil"
	"net/http"
)

func setJsonHeaders(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")
}

func setCORSHeaders(resp http.ResponseWriter, allowOrigin string) {
	resp.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	resp.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	resp.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,user,authorization")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
}

//CORSProvider can be used as middleware.
type CORSProvider struct {
	AllowOrigin string
}

func (cp *CORSProvider) ServeHTTP(resp http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	if cp.AllowOrigin == "" {
		cp.AllowOrigin = "*"
	}
	setCORSHeaders(resp, cp.AllowOrigin)
	next(resp, req)
}

//Marshal any Marshalable type
func Marshal(resp http.ResponseWriter, statusCode int, result interface{}) {
	marshal(resp, statusCode, result, true)
}

func marshal(resp http.ResponseWriter, statusCode int, result interface{}, writeHeader bool) {
	b, err := json.Marshal(result)
	if err != nil {
		logrus.Warnf("Marshal result error: %+v", err)
		b2, err := json.Marshal(err.Error())
		if err != nil {
			//if marhsalling error-only results in another error, don't use JSON content. Just text/plain:
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		} else {
			setJsonHeaders(resp)
			if writeHeader {
				resp.WriteHeader(http.StatusInternalServerError)
			}
			fmt.Fprint(resp, string(b2))
		}
		return
	}
	setJsonHeaders(resp)
	if writeHeader {
		resp.WriteHeader(statusCode)
	}
	fmt.Fprint(resp, string(b))
}

//helper to load a provided pointer from json
func LoadFromJson(req io.Reader, i interface{}) error {
	jsonBlob, err := ioutil.ReadAll(req)
	if err != nil {
		//using warnings because it's usually a client error
		logrus.Warnf("Read body error: %+v", err)
		return err
	}
	err = json.Unmarshal(jsonBlob, i)
	if err != nil {
		logrus.Warnf("Unmarshal error: %+v", err)
		logrus.Warnf("Blob: %+v", string(jsonBlob))
		return err
	}
	return err
}

type ErrStruct struct {
	Message string `json:",omitempty"`
	Code    int    `json:",omitempty"`
}

func ErrorCalc(resp http.ResponseWriter, err error) {
	statusMessage, code := GetErrorHttpStatus(err)
	Error(resp, statusMessage, code)
}

func Error(resp http.ResponseWriter, err string, code int) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(code)
	if code == http.StatusInternalServerError {
		logrus.Errorf("Server error: %s", err)
	} else if code > 199 && code < 300 {
		logrus.Debugf("Returning OK code %d: %s", code, err)
	} else {
		logrus.Warnf("Returning error code %d: %s", code, err)
	}
	marshal(resp, code, ErrStruct{err, code}, false)
}
