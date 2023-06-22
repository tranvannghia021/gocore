package repositories

import "time"

type StructureAuth struct {
	ClientId            string `json:"client_id"`
	RedirectUri         string `json:"redirect_uri"`
	Scope               string `json:"scope"`
	ResponseType        string `json:"response_type"`
	State               string `json:"state"`
	CodeChallenge       string `json:"code_challenge,omitempty"`
	CodeChallengeMethod string `json:"code_challenge_method,omitempty"`
}

type ConfigSocial struct {
	Separator   string
	UsePKCE     bool
	Scopes      []string
	Paramerters map[string]string
}

type PayloadGenerate struct {
	Ip           string    `json:"ip,omitempty"`
	Platform     string    `json:"platform,omitempty"`
	CodeVerifier string    `json:"code_verifier,omitempty"`
	ID           int       `json:"id ,omitempty"`
	Email        string    `json:"email,omitempty"`
	CreateAt     time.Time `json:"create_at,omitempty"`
}
