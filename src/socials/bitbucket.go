package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"net/url"
	"time"
)

var bitbucket = "bitbucket"

type sBitbucket struct {
	http *service.SHttpRequest
}

var scopeBb []string

type profileBb struct {
	DisplayName string `json:"display_name"`
	Links       struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		Repositories struct {
			Href string `json:"href"`
		} `json:"repositories"`
		Snippets struct {
			Href string `json:"href"`
		} `json:"snippets"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Hooks struct {
			Href string `json:"href"`
		} `json:"hooks"`
	} `json:"links"`
	CreatedOn     time.Time `json:"created_on"`
	Type          string    `json:"type"`
	UUID          string    `json:"uuid"`
	Has2FaEnabled any       `json:"has_2fa_enabled"`
	Username      string    `json:"username"`
	IsStaff       bool      `json:"is_staff"`
	AccountID     string    `json:"account_id"`
	Nickname      string    `json:"nickname"`
	AccountStatus string    `json:"account_status"`
	Location      string    `json:"location"`
}
type emailBb struct {
	Values []struct {
		Type  string `json:"type"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
		Email       string `json:"email"`
		IsPrimary   bool   `json:"is_primary"`
		IsConfirmed bool   `json:"is_confirmed"`
	} `json:"values"`
	Pagelen int `json:"pagelen"`
	Size    int `json:"size"`
	Page    int `json:"page"`
}

func (s *sBitbucket) loadConfig() {
	coreConfig.Separator = ","
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"email",
	}, scopeBb...))
	urlAuth = "https://bitbucket.org/site/oauth2/authorize"
	s.http = service.NewHttpRequest()
}

func (s *sBitbucket) getToken(code string) vars.ResReq {
	s.http.FormData, _ = buildPayloadToken(code, true)
	s.http.Url = "https://bitbucket.org/site/oauth2/access_token"
	s.http.Headers = headerAuthBasic()
	return s.http.PostFormDataRequest()
}

func (s *sBitbucket) profile(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	param := query.Encode()
	s.http.Url = fmt.Sprintf("%s/%s/user?%s", vars.EndPoint, vars.Version, param)
	result := s.http.GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	resultEmail := s.getEmail(param)
	if !resultEmail.Status {
		helpers.CheckNilErr(resultEmail.Error)
		return repositories.Core{}
	}
	var email emailBb
	var profile profileBb

	_ = json.Unmarshal(result.Data, &profile)
	_ = json.Unmarshal(resultEmail.Data, &email)
	return repositories.Core{
		InternalId:    profile.AccountID,
		Platform:      bitbucket,
		Email:         email.Values[0].Email,
		EmailVerifyAt: time.Time{},
		FirstName:     profile.DisplayName,
		Avatar:        profile.Links.Avatar.Href,
		Status:        true,
		BirthDay:      time.Time{},
		Address:       profile.Location,
	}
}

func AddScopeBitbucket(scope []string) {
	scopeBb = scope
}
func (s *sBitbucket) getEmail(param string) vars.ResReq {
	s.http.Url = fmt.Sprintf("%s/%s/user/emails?%s", vars.EndPoint, vars.Version, param)
	return s.http.GetRequest()
}
