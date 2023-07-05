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

var line = "line"

type sLine struct {
	http *service.SHttpRequest
}

var scopeL []string
var idToken string

type profileL struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
}

func (s *sLine) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.UsePKCE = true
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"profile",
		"openid",
	}, scopeL...))
	urlAuth = "https://access.line.me/oauth2/v2.1/authorize"
	s.http = service.NewHttpRequest()
}

func (s *sLine) getToken(code string) vars.ResReq {
	s.http.FormData, _ = buildPayloadToken(code, true)
	s.http.Url = fmt.Sprintf("%s/oauth2/%s/token", vars.EndPoint, vars.Version)
	return s.http.PostFormDataRequest()
}

func (s *sLine) profile(token string) repositories.Core {
	verify := s.verifyToken(idToken)
	if !verify.Status {
		helpers.CheckNilErr(verify.Error)
		return repositories.Core{}
	}
	s.http.Url = fmt.Sprintf("%s/oauth2/%s/userinfo", vars.EndPoint, vars.Version)
	result := s.http.SetAuth(token).GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var resVerify verifyTokenResL
	var profile profileL
	_ = json.Unmarshal(verify.Data, &resVerify)
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    profile.Sub,
		Platform:      line,
		Email:         resVerify.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name,
		Avatar:        resVerify.Picture,
		Status:        true,
	}

}

type verifyTokenResL struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     int      `json:"exp"`
	Iat     int      `json:"iat"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Picture string   `json:"picture"`
}

func (g sLine) verifyToken(idToken string) vars.ResReq {
	body := url.Values{}
	body.Add("id_token", idToken)
	body.Add("client_id", vars.ClientId)
	g.http.Url = fmt.Sprintf("%s/oauth2/%s/verify", vars.EndPoint, vars.Version)
	g.http.FormData = body
	return g.http.PostFormDataRequest()
}

func AddScopeLine(scope []string) {
	scopeL = scope
}
