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

var facebook string = "facebook"
var defaultScopeFb = []string{
	"public_profile",
	"email"}
var scopeFb []string
var defaultFieldsFb = []string{
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
}
var fieldFb []string

func init() {
	AddScopeFaceBook(scopeFb)
	AddFieldFacebook(fieldFb)
	vars.CallConfig[facebook] = loadConfigFb
	vars.PLatFormToken[facebook] = getTokenFb
	vars.PLatFormProfile[facebook] = profileFb
}

func AddScopeFaceBook(scope []string) {
	scopeFb = scope
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append(defaultScopeFb, scope...))
}

func AddFieldFacebook(fields []string) {
	fieldFb = fields
	coreConfig.Fields = helpers.RemoveDuplicateStr(append(defaultFieldsFb, fields...))
}

var getTokenFb = func(code string) vars.ResReq {
	url := fmt.Sprintf("%s/%s/oauth/access_token?%s", vars.EndPoint, vars.Version, buildPayloadToken(code))
	return service.GetRequest(url, make(map[string]string))
}

var profileFb = func(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	results := service.GetRequest(fmt.Sprintf("%s/%s/me?%s&fields=%s",
		vars.EndPoint,
		vars.Version, query.Encode(), strings.Join(coreConfig.Fields, ",")), nil)
	if !results.Status {
		helpers.CheckNilErr(results.Error)
	}
	var profile ProfileFB
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

var loadConfigFb = func() {
	coreConfig.UsePKCE = false
	coreConfig.Separator = ","
	urlAuth = fmt.Sprintf("https://www.facebook.com/%s/dialog/oauth", vars.Version)
	parameters["display"] = "popup"
}

type ProfileFB struct {
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
