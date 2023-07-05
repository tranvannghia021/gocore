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
	http *service.SHttpRequest
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

func (s *sTwitter) loadConfig() {
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
	s.http = service.NewHttpRequest()
}

func (s *sTwitter) getToken(code string) vars.ResReq {
	s.http.FormData, _ = buildPayloadToken(code, true)
	s.http.Url = fmt.Sprintf("%s/%s/oauth2/token", vars.EndPoint, vars.Version)
	return s.http.PostFormDataRequest()
}

func (s *sTwitter) profile(token string) repositories.Core {
	param := url.Values{}
	param.Add("user.fields", strings.Join(coreConfig.Fields, ","))
	s.http.Url = fmt.Sprintf("%s/%s/users/me?%s", vars.EndPoint, vars.Version, param.Encode())
	result := s.http.SetAuth(token).GetRequest()
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
