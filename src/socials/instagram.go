package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"net/url"
	"strings"
	"time"
)

var instagram = "instagram"

var scopeInS []string

var fieldInS []string

type sInstagram struct {
}
type profileInS struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (s sInstagram) loadConfig() {
	coreConfig.UsePKCE = false
	coreConfig.Separator = ","
	urlAuth = "https://api.instagram.com/oauth/authorize"
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"email",
		"user_profile",
	}, scopeInS...))
	coreConfig.Fields = helpers.RemoveDuplicateStr(append([]string{
		"id",
		"username",
	}, fieldInS...))
}

func (s sInstagram) getToken(code string) vars.ResReq {
	data, _ := buildPayloadToken(code, true)
	url := "https://api.instagram.com/oauth/access_token"
	return service.PostFormDataRequest(url, nil, data)
}

func (s sInstagram) profile(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	query.Add("fields", strings.Join(fieldInS, ","))
	results := service.GetRequest(fmt.Sprintf("%s/me?%s",
		vars.EndPoint, query.Encode()), nil)
	if !results.Status {
		helpers.CheckNilErr(results.Error)
		return repositories.Core{}
	}
	var profile profileInS
	_ = json.Unmarshal(results.Data, &profile)
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      instagram,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Username,
		Status:        true,
	}
}

func AddScopeInstagram(scope []string) {
	scopeInS = scope
}

func AddFieldInstagram(fields []string) {
	fieldInS = fields
}
