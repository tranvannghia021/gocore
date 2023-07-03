package middlewares

import (
	errors2 "errors"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/response"
	"net/http"
)

func Refresh(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state := r.Header.Get("Authorization")
		var errors []string
		if state == "" {
			errors = append(errors, "Authentication failed!")
		} else {
			data, expire := helpers.DecodeJWT(state, true)
			if expire {
				errors = append(errors, "Signature is invalid!")
			}
			r.Header.Set("platform", data.Platform)
			r.Header.Set("code_verifier", data.CodeVerifier)
			r.Header.Set("ip", data.Ip)
			r.Header.Set("email", data.Email)
			r.Header.Set("id", data.ID.String())

		}

		if len(errors) > 0 {
			response.Response(nil, w, true, errors2.New(errors[0]))
			return
		}

		next.ServeHTTP(w, r)
	})
}
