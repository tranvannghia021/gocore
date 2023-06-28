package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"time"
)

var google = "google"
var scopeGg []string

type sGoogle struct {
}
type profileGg struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func (sGoogle) loadConfig() {
	coreConfig.UsePKCE = false
	coreConfig.Separator = " "
	urlAuth = fmt.Sprintf("https://accounts.google.com/o/oauth2/%s/auth", vars.Version)
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/user.addresses.read",
		"https://www.googleapis.com/auth/user.birthday.read",
		"https://www.googleapis.com/auth/user.emails.read",
		"https://www.googleapis.com/auth/user.gender.read",
		"https://www.googleapis.com/auth/user.organization.read",
		"https://www.googleapis.com/auth/user.phonenumbers.read",
		"https://www.googleapis.com/auth/userinfo.profile",
	}, scopeGg...))
}

func AddScopeGoogle(scope []string) {
	scopeGg = scope
}

func (g sGoogle) getToken(code string) vars.ResReq {
	_, data := buildPayloadToken(code, false)
	url := fmt.Sprintf("https://oauth2.%s/token", vars.EndPoint)
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"
	return service.PostRequest(url, headers, data)
}

func (g sGoogle) profile(token string) repositories.Core {
	results := service.GetRequest(fmt.Sprintf("https://www.%s/oauth2/%s/userinfo?alt=json&access_token=%s",
		vars.EndPoint,
		vars.Version, token), map[string]string{})
	if !results.Status {
		helpers.CheckNilErr(results.Error)
	}
	var profile profileGg
	_ = json.Unmarshal(results.Data, &profile)
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      google,
		Email:         profile.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name,
		Avatar:        profile.Picture,
		Status:        true,
	}
}
