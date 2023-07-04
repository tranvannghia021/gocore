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

type sFacebook struct {
}

var facebook = "facebook"
var scopeFb []string
var fieldFb []string

type profileFB struct {
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

func (sFacebook) loadConfig() {
	coreConfig.UsePKCE = false
	coreConfig.Separator = ","
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"public_profile",
		"email",
	}, scopeFb...))
	coreConfig.Fields = helpers.RemoveDuplicateStr(append([]string{
		"id",
		"name",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"hometown",
		"location",
		"picture",
	}, fieldFb...))
	urlAuth = fmt.Sprintf("https://www.facebook.com/%s/dialog/oauth", vars.Version)
	parameters["display"] = "popup"
}

func (sFacebook) getToken(code string) vars.ResReq {
	data, _ := buildPayloadToken(code, true)
	url := fmt.Sprintf("%s/%s/oauth/access_token?%s", vars.EndPoint, vars.Version, data.Encode())
	return service.GetRequest(url, make(map[string]string))
}

func AddScopeFaceBook(scope []string) {
	scopeFb = scope
}

func AddFieldFacebook(fields []string) {
	fieldFb = fields
}
func (sFacebook) profile(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	results := service.GetRequest(fmt.Sprintf("%s/%s/me?%s&fields=%s",
		vars.EndPoint,
		vars.Version, query.Encode(), strings.Join(coreConfig.Fields, ",")), nil)
	if !results.Status {
		helpers.CheckNilErr(results.Error)
		return repositories.Core{}
	}
	var profile profileFB
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
