package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/config"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/service"
	"github.com/tranvannghia021/gocore/src/repositories"
	"net/url"
	"strings"
	"time"
)

var facebook string = "facebook"

func init() {
	config.CallConfig[facebook] = loadConfigFb
	config.PLatFormToken[facebook] = getTokenFb
	config.PLatFormProfile[facebook] = profileFb
}

func AddScopeFaceBook(scope []string) {
	coreConfig.Scopes = scope
}

func AddFieldFacebook(fields []string) {
	coreConfig.Fields = fields
}

var getTokenFb = func(code string) config.ResReq {
	url := fmt.Sprintf("%s/%s/oauth/access_token?%s", config.EndPoint, config.Version, buildPayloadToken(code))
	return service.GetRequest(url, make(map[string]string))
}

var profileFb = func(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	results := service.GetRequest(fmt.Sprintf("%s/%s/me?%s&fields=%s",
		config.EndPoint,
		config.Version, query.Encode(), strings.Join(coreConfig.Fields, ",")), nil)
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
	urlAuth = fmt.Sprintf("https://www.facebook.com/%s/dialog/oauth", config.Version)
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
