package gocore

import (
	"github.com/tranvannghia021/gocore/src"
	"net/http"
)

var core = src.NewHandler()
var manager = src.NewManager()

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	core.GenerateUrl(w, r)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	core.AuthHandle(w, r)
}
func SignIn(w http.ResponseWriter, r *http.Request) {
	manager.SignIn(w, r)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	manager.SignUp(w, r)
}
func Update(w http.ResponseWriter, r *http.Request) {
	manager.Update(w, r)
}
func Me(w http.ResponseWriter, r *http.Request) {
	manager.Me(w, r)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	manager.Delete(w, r)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	manager.Refresh(w, r)
}

func Resend(w http.ResponseWriter, r *http.Request) {
	manager.Resend(w, r)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	manager.Verify(w, r)
}

func Success(w http.ResponseWriter, r *http.Request) {
	manager.Success(w, r)
}
