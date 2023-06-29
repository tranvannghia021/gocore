package gocore

import (
	"github.com/tranvannghia021/gocore/src"
	"net/http"
)

var core = src.Shandler{}

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	core.GenerateUrl(w, r)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	core.AuthHandle(w, r)
}
