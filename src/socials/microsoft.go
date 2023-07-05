package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"strings"
	"time"
)

var microsoft = "microsoft"

type sMicrosoft struct {
	http *service.SHttpRequest
}

var scopeMs []string

type profileMs struct {
	OdataContext      string `json:"@odata.context"`
	BusinessPhones    []any  `json:"businessPhones"`
	DisplayName       string `json:"displayName"`
	GivenName         string `json:"givenName"`
	JobTitle          any    `json:"jobTitle"`
	Mail              string `json:"mail"`
	MobilePhone       any    `json:"mobilePhone"`
	OfficeLocation    any    `json:"officeLocation"`
	PreferredLanguage any    `json:"preferredLanguage"`
	Surname           string `json:"surname"`
	UserPrincipalName string `json:"userPrincipalName"`
	ID                string `json:"id"`
}

func (s *sMicrosoft) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"user.read",
		"offline_access",
		"mail.read",
	}, scopeMs...))
	urlAuth = fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/authorize", vars.Tenant)
	parameters["response_mode"] = "query"
	s.http = service.NewHttpRequest()
}

func (s *sMicrosoft) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	body.Add("scope", strings.Join(coreConfig.Scopes, coreConfig.Separator))
	s.http.Url = fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", vars.Tenant)
	s.http.FormData = body
	return s.http.PostFormDataRequest()
}

func (s *sMicrosoft) profile(token string) repositories.Core {
	s.http.Url = fmt.Sprintf("%s/%s/me", vars.EndPoint, vars.Version)
	result := s.http.SetAuth(token).GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var profile profileMs
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      microsoft,
		Email:         profile.Mail,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.DisplayName,
		Status:        true,
	}
}

func AddScopeMicrosoft(scope []string) {
	scopeMs = scope
}
