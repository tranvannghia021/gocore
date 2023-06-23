package src

import (
	"errors"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/socials"
	"net/http"
)

func GenerateUrl(platform string) string {
	return socials.New(platform).Generate()
}

func AuthHandle(r *http.Request) {
	if r.URL.Query().Get("error") != "" {
		helpers.CheckNilErr(errors.New("Accept denied!"))
		return
	}

	if r.URL.Query().Get("errors") != "" {
		helpers.CheckNilErr(errors.New("Accept denied!"))
		return
	}
	socials.New(r.Header.Get("platform")).Auth(r.URL.Query().Get("code"), r)
}
