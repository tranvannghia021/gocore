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

func (s sMicrosoft) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"user.read",
		"offline_access",
		"mail.read",
	}, scopeMs...))
	urlAuth = fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/authorize", vars.Tenant)
	parameters["response_mode"] = "query"
}

func (s sMicrosoft) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	body.Add("scope", strings.Join(coreConfig.Scopes, coreConfig.Separator))
	return service.PostFormDataRequest(fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", vars.Tenant), nil, body)
}

func (s sMicrosoft) profile(token string) repositories.Core {
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	result := service.GetRequest(fmt.Sprintf("%s/%s/me", vars.EndPoint, vars.Version), headers)
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
