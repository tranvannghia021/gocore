package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	errors2 "errors"
	"github.com/tranvannghia021/gocore/src/response"
	"net/http"
	"net/url"
	"os"
)

func VerifyHmac(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("platform") == "shopify" && !checkHmac(r) {
			response.Response(nil, w, true, errors2.New("Authentication is invalid"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func checkHmac(r *http.Request) bool {
	params := url.Values{}
	params.Add("code", r.URL.Query().Get("code"))
	params.Add("host", r.URL.Query().Get("host"))
	params.Add("shop", r.URL.Query().Get("shop"))
	params.Add("timestamp", r.URL.Query().Get("timestamp"))
	params.Add("state", r.URL.Query().Get("state"))
	secret, _ := os.LookupEnv("SHOPIFY_CLIENT_SECRET")
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(params.Encode()))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha == r.URL.Query().Get("hmac")
}
