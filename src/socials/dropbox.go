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

var dropbox = "dropbox"

type sDropbox struct {
	http *service.SHttpRequest
}

var scopeDb []string

type profileDb struct {
	AccountID string `json:"account_id"`
	Name      struct {
		GivenName       string `json:"given_name"`
		Surname         string `json:"surname"`
		FamiliarName    string `json:"familiar_name"`
		DisplayName     string `json:"display_name"`
		AbbreviatedName string `json:"abbreviated_name"`
	} `json:"name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Disabled      bool   `json:"disabled"`
	Country       string `json:"country"`
	Locale        string `json:"locale"`
	ReferralLink  string `json:"referral_link"`
	IsPaired      bool   `json:"is_paired"`
	AccountType   struct {
		Tag string `json:".tag"`
	} `json:"account_type"`
	RootInfo struct {
		Tag             string `json:".tag"`
		RootNamespaceID string `json:"root_namespace_id"`
		HomeNamespaceID string `json:"home_namespace_id"`
	} `json:"root_info"`
}

func (s *sDropbox) loadConfig() {
	coreConfig.UsePKCE = true
	coreConfig.Separator = ","
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"account_info.read",
	}, scopeDb...))
	urlAuth = "https://www.dropbox.com/oauth2/authorize"
	s.http = service.NewHttpRequest()
}

func (s *sDropbox) getToken(code string) vars.ResReq {
	s.http.FormData, _ = buildPayloadToken(code, true)
	s.http.Url = fmt.Sprintf("%s/oauth2/token", vars.EndPoint)
	return s.http.PostFormDataRequest()
}

func (s *sDropbox) profile(token string) repositories.Core {
	s.http.Url = fmt.Sprintf("%s/%s/users/get_current_account", vars.EndPoint, vars.Version)
	s.http.PostRequest()
	result := s.http.SetAuth(token).PostRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var profile profileDb
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    profile.AccountID,
		Platform:      dropbox,
		Email:         profile.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name.FamiliarName,
		LastName:      profile.Name.Surname,
		Status:        true,
	}
}

func AddScopeDropbox(scope []string) {
	scopeDb = scope
}
