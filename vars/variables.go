package vars

import (
	"github.com/go-redis/redis"
	"github.com/tranvannghia021/gocore/src/repositories"
	"gorm.io/gorm"
	"time"
)

var Payload = PayloadGenerate{}
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

var Connection *gorm.DB
var Redis *redis.Client

type ResReq struct {
	Status bool
	Data   []byte
	Error  error
}
type ConfigCore struct {
	Database struct {
		TableName string            `json:"table_name"`
		Fields    map[string]string `json:"fields"`
	} `json:"database"`
}
type PayloadGenerate struct {
	Ip           string
	Platform     string
	CodeVerifier string
	ID           int
	Email        string
	CreateAt     time.Time
}
