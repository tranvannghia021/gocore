package gocore

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tranvannghia021/gocore/config"
	"github.com/tranvannghia021/gocore/src"
	"github.com/tranvannghia021/gocore/src/response"
	"github.com/tranvannghia021/gocore/src/socials"
	"net/http"
	"os"
)

func init() {
	godotenv.Load()
	config.ConnectDB()
	config.ConnectCache()
	LoadCoreConfig()
}

type StyleRpPusher struct {
	Url    string `json:"url"`
	Pusher struct {
		Channel string `json:"channel"`
		Event   string `json:"event"`
	} `json:"pusher"`
}

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	platform := mux.Vars(r)["platform"]
	config.Payload.Ip = r.RemoteAddr
	config.Payload.Platform = platform
	url := src.GenerateUrl(platform)
	channel, _ := os.LookupEnv("CHANNEL_NAME")
	event, _ := os.LookupEnv("EVENT_NAME")
	response.Response(&StyleRpPusher{
		Url: url,
		Pusher: struct {
			Channel string `json:"channel"`
			Event   string `json:"event"`
		}{
			Channel: channel,
			Event:   event + "_" + config.Payload.Ip,
		},
	}, w, false, nil)
}

func Auth(w http.ResponseWriter, r *http.Request) {
	src.AuthHandle(r)
}

func Migrate(w http.ResponseWriter, r *http.Request) {

}

func Rollback(w http.ResponseWriter, r *http.Request) {
	config.RollbackMigrate(config.Connection)
	json.NewEncoder(w).Encode("Ok")
}

func LoadCoreConfig() {
	socials.AddScopeFaceBook([]string{
		"public_profile",
		"email",
	})

	socials.AddFieldFacebook([]string{
		"id",
		"name",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"hometown",
		"location",
		"picture",
	})
}
