package repositories

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
	Fields      []string
}
