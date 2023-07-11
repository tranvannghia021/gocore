package gocore

import (
	"net/http"
)

func GenerateUrlRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	core.GenerateUrl(w, r)
}

func AuthRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	core.AuthHandle(w, r)
}
func SignInRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.SignIn(w, r)
}

func SignUpRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.SignUp(w, r)
}
func UpdateRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Update(w, r)
}
func MeRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Me(w, r)
}
func DeleteRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Delete(w, r)
}

func RefreshRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Refresh(w, r)
}

func ResendRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Resend(w, r)
}

func VerifyRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Verify(w, r)
}

func SuccessRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Success(w, r)
}

func ForgotRunTime(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	manager.Forgot(w, r)
}
