package repositories

import (
	"github.com/google/uuid"
	"time"
)

type Core struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey:type:uuid"`
	InternalId    string    `json:"internal_id,omitempty" gorm:"unique"`
	Platform      string    `json:"platform,omitempty"`
	Email         string    `json:"email,omitempty"`
	EmailVerifyAt time.Time `json:"email_verify_at,omitempty"`
	Password      string    `json:"password,omitempty"`
	FirstName     string    `json:"first_name,omitempty"`
	LastName      string    `json:"last_name,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Gender        string    `json:"gender,omitempty"`
	Status        bool      `json:"status,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	BirthDay      time.Time `json:"birthday,omitempty"`
	Address       string    `json:"address,omitempty"`
	RefreshToken  string    `json:"refresh_token,omitempty"`
	AccessToken   string    `json:"access_token,omitempty"`
	ExpireToken   time.Time `json:"expire_token,omitempty"`
	IsDisconnect  bool      `json:"is_disconnect,omitempty"`
	Domain        string    `json:"domain,omitempty"`
	RawDomain     string    `json:"raw_domain,omitempty"`
}
