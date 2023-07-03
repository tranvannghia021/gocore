package vars

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	mail "github.com/xhit/go-simple-mail/v2"
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

var Connection *gorm.DB
var Redis *redis.Client
var Mail *mail.SMTPClient

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
	ID           uuid.UUID
	Email        string
	CreateAt     time.Time
	Domain       string
}
