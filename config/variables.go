package config

import "github.com/tranvannghia021/gocore/src/repositories"

var (
	AppId        string
	ClientId     string
	ClientSecret string
	EndPoint     string
	Version      string
	RedirectUri  string
	Tenant       string
)
var CallConfig = make(map[string]func())
var PLatFormToken = make(map[string]func(code string) ResReq)
var PLatFormProfile = make(map[string]func(code string) repositories.Core)
var Payload = repositories.PayloadGenerate{}

type ResReq struct {
	Status bool
	Data   []byte
	Error  error
}
