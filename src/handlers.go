package src

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/response"
	"github.com/tranvannghia021/gocore/src/socials"
	"github.com/tranvannghia021/gocore/vars"
	"net/http"
	"os"
)

type IHandler interface {
	GenerateUrl(w http.ResponseWriter, r *http.Request)
	AuthHandle(w http.ResponseWriter, r *http.Request)
}
type Shandler struct {
}
type styleRpPusher struct {
	Url    string `json:"url"`
	Pusher struct {
		Channel string `json:"channel"`
		Event   string `json:"event"`
	} `json:"pusher"`
}
type sShopify struct {
	Domain string `json:"domain"`
}

func (Shandler) GenerateUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vars.Payload.Ip = helpers.Md5(r.RemoteAddr)
	vars.Payload.Platform = params["platform"]
	var s sShopify
	_ = json.NewDecoder(r.Body).Decode(&s)
	vars.Payload.Domain = s.Domain
	channel, _ := os.LookupEnv("CHANNEL_NAME")
	event, _ := os.LookupEnv("EVENT_NAME")
	response.Response(&styleRpPusher{
		Url: socials.New().Generate(),
		Pusher: struct {
			Channel string `json:"channel"`
			Event   string `json:"event"`
		}{
			Channel: channel,
			Event:   event + "_" + vars.Payload.Ip,
		},
	}, w, false, nil)
}

func (Shandler) AuthHandle(w http.ResponseWriter, r *http.Request) {
	vars.Payload.Platform = r.Header.Get("platform")
	params := r.URL.Query()
	vars.Payload.Domain = params.Get("shop")
	if params.Get("error") != "" || params.Get("errors") != "" {
		helpers.CheckNilErr(errors.New("Accept denied!"))
		return
	}
	socials.New().Auth(r)
	fontEndUrl, _ := os.LookupEnv("FONT_END_URL")
	http.Redirect(w, r, fontEndUrl, 302)
}
