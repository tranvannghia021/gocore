package gocore

import (
	"github.com/tranvannghia021/gocore/src"
	"net/http"
)

func new() src.IHandler {
	return src.Shandler{}
}
func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	new().GenerateUrl(w, r)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	new().AuthHandle(w, r)
}
