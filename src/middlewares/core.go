package middlewares

import (
	"github.com/tranvannghia021/gocore/vars"
	"net/http"
)

func Core(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars.Wh = w // exception
		//todo
		next.ServeHTTP(w, r)
	})
}
