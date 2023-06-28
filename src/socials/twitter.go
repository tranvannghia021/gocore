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

var twitter = "twitter"

type sTwitter struct {
}

var (
	scopeTt  []string
	fieldsTt []string
)

type profileTt struct {
	Data struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Username        string `json:"username"`
		ProfileImageURL string `json:"profile_image_url"`
	} `json:"data"`
}

func (s sTwitter) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.UsePKCE = true
	urlAuth = "https://twitter.com/i/oauth2/authorize"
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"users.read",
		"tweet.read",
	}, scopeTt...)) //id,name,profile_image_url,url,username
	coreConfig.Fields = helpers.RemoveDuplicateStr(append([]string{
		"id",
		"name",
		"profile_image_url",
		"url",
		"username",
	}, fieldsTt...))
}

func (s sTwitter) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	return service.PostFormDataRequest(fmt.Sprintf("%s/%s/oauth2/token", vars.EndPoint, vars.Version), nil, body)
}

func (s sTwitter) profile(token string) repositories.Core {
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	param := url.Values{}
	param.Add("user.fields", strings.Join(coreConfig.Fields, ","))
	result := service.GetRequest(fmt.Sprintf("%s/%s/users/me?%s", vars.EndPoint, vars.Version, param.Encode()), headers)
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var profile profileTt
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    profile.Data.ID,
		Platform:      twitter,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Data.Name,
		Avatar:        profile.Data.ProfileImageURL,
		Status:        true,
	}
}

func AddScopeTwitter(scope []string) {
	scopeTt = scope
}

func AddFieldTwitter(fields []string) {
	fieldsTt = fields
}
