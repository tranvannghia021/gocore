package middlewares

import (
	errors2 "errors"
	"github.com/google/uuid"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	sql2 "github.com/tranvannghia021/gocore/src/repositories/sql"
	"github.com/tranvannghia021/gocore/src/response"
	"github.com/tranvannghia021/gocore/vars"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		state := r.Header.Get("Authorization")
		var errors []string
		if state == "" {
			errors = append(errors, "Authentication failed!")
		} else {
			data, expire := helpers.DecodeJWT(state, false)
			if expire {
				errors = append(errors, "Signature is invalid!")
			}
			r.Header.Set("platform", data.Platform)
			r.Header.Set("code_verifier", data.CodeVerifier)
			r.Header.Set("ip", data.Ip)
			r.Header.Set("email", data.Email)
			globalVariantUser(data.ID)
			r.Header.Set("id", data.ID.String())

		}

		if len(errors) > 0 {
			response.Response(nil, w, true, errors2.New(errors[0]))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func globalVariantUser(id uuid.UUID) {
	var core = repositories.Core{ID: id}
	sql := new(sql2.Suser)
	sql.ModelBase = &core
	result := sql.First()
	if result.Status {
		vars.User = &core
	}

}
