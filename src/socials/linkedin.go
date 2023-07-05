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
	"time"
)

var linkedin = "linkedin"

type sLinkedin struct {
	http *service.SHttpRequest
}

var scopeLk []string

type profileLk struct {
	FirstName struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"firstName"`
	LastName struct {
		Localized struct {
			EnUS string `json:"en_US"`
		} `json:"localized"`
		PreferredLocale struct {
			Country  string `json:"country"`
			Language string `json:"language"`
		} `json:"preferredLocale"`
	} `json:"lastName"`
	ID string `json:"id"`
}

func (s *sLinkedin) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"r_liteprofile",
		"r_emailaddress",
	}, scopeLk...))
	s.http = service.NewHttpRequest()
	urlAuth = fmt.Sprintf("https://www.linkedin.com/oauth/%s/authorization", vars.Version)
}

func (s *sLinkedin) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	s.http.FormData = body
	s.http.Url = fmt.Sprintf("https://www.linkedin.com/oauth/%s/accessToken?%s", vars.Version, body.Encode())
	return s.http.PostFormDataRequest()
}

func (s *sLinkedin) profile(token string) repositories.Core {
	param := url.Values{}
	param.Add("projection", "(id,firstName,lastName,profilePicture(displayImage~:playableStreams))")
	s.http.Url = fmt.Sprintf("%s/%s/me?%s", vars.EndPoint, vars.Version, param.Encode())
	result := s.http.SetAuth(token).GetRequest()
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	resultEmail := s.getEmailLinkedin(token)
	if !resultEmail.Status {
		helpers.CheckNilErr(resultEmail.Error)
		return repositories.Core{}
	}
	var profile profileLk
	var email emailLk
	_ = json.Unmarshal(result.Data, &profile)
	_ = json.Unmarshal(resultEmail.Data, &email)
	log.Println(string(result.Data))
	return repositories.Core{
		InternalId:    profile.ID,
		Platform:      linkedin,
		Email:         email.Elements[0].Handle.EmailAddress,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.FirstName.Localized.EnUS,
		LastName:      profile.LastName.Localized.EnUS,
		Status:        true,
	}
}

type emailLk struct {
	Elements []struct {
		Handle struct {
			EmailAddress string `json:"emailAddress"`
		} `json:"handle~"`
		Handle0 string `json:"handle"`
	} `json:"elements"`
}

func (s *sLinkedin) getEmailLinkedin(token string) vars.ResReq {
	query := url.Values{}
	query.Add("q", "members")
	query.Add("projection", "(elements*(handle~))")
	s.http.Url = fmt.Sprintf("%s/%s/emailAddress?%s", vars.EndPoint, vars.Version, query.Encode())
	return s.http.SetAuth(token).GetRequest()
}

func AddScopeLinkedin(scope []string) {
	scopeLk = scope
}
