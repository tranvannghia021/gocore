package gocore

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tranvannghia021/gocore/config"
	"github.com/tranvannghia021/gocore/src"
	"net/http"
)

var (
	AppId        string
	ClientId     string
	ClientSecret string
	EndPoint     string
	Version      string
	RedirectUri  string
)

func init() {
	godotenv.Load()
	config.ConnectDB()
	config.ConnectCache()
}

func GenerateUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	platform := mux.Vars(r)["platform"]
	url := src.Social(platform)
	json.NewEncoder(w).Encode(url)
}
