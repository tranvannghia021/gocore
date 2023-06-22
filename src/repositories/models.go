package repositories

import "time"

type Core struct {
	Id            int         `json:"id"`
	InternalId    string      `json:"internal_id"`
	Platform      string      `json:"platform"`
	Email         string      `json:"email"`
	EmailVerifyAt time.Time   `json:"email_verify_at"`
	Password      string      `json:"password"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Avatar        string      `json:"avatar"`
	Gender        string      `json:"gender"`
	Status        bool        `json:"status"`
	Phone         string      `json:"phone"`
	BirthDay      time.Time   `json:"birthday"`
	Address       string      `json:"address"`
	RefreshToken  string      `json:"refresh_token"`
	AccessToken   string      `json:"access_token"`
	ExpireToken   time.Time   `json:"expire_token"`
	IsDisconnect  bool        `json:"is_disconnect"`
	Domain        string      `json:"domain"`
	RawDomain     string      `json:"raw_domain"`
	Settings      interface{} `json:"settings"`
}
