package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"log"
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
	ID        string `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Picture   struct {
		Data struct {
			Height       int    `json:"height"`
			IsSilhouette bool   `json:"is_silhouette"`
			URL          string `json:"url"`
			Width        int    `json:"width"`
		} `json:"data"`
	} `json:"picture"`
}

func (s sInstagram) loadConfig() {
	coreConfig.UsePKCE = false
	coreConfig.Separator = ","
	urlAuth = "https://api.instagram.com/oauth/authorize"
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"email",
		"public_profile",
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
	results := service.GetRequest(fmt.Sprintf("%s/me?",
		vars.EndPoint, query.Encode()), nil)
	if !results.Status {
		helpers.CheckNilErr(results.Error)
	}
	log.Println(string(results.Data))
	var profile profileInS
	_ = json.Unmarshal(results.Data, &profile)
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      facebook,
		Email:         profile.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.FirstName,
		LastName:      profile.LastName,
		Avatar:        profile.Picture.Data.URL,
		Status:        true,
	}
}

func AddScopeInstagram(scope []string) {
	scopeInS = scope
}

func AddFieldInstagram(fields []string) {
	fieldInS = fields
}
