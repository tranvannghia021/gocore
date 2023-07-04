package socials

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/tranvannghia021/gocore/config"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/repositories/sql"
	"github.com/tranvannghia021/gocore/vars"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type social interface {
	Generate() string
	Auth(r *http.Request)
	buildLinkAuth(state string) string
	getUrlAuth() string
	formatScope() string
	getCodeChallenge() string
}
type iCore interface {
	loadConfig()
	getToken(code string) vars.ResReq
	profile(token string) repositories.Core
}

var codeVerifier string
var urlAuth string
var parameters = make(map[string]string)
var coreConfig repositories.ConfigSocial

type resToken struct {
	AccessToken          string `json:"access_token,omitempty"`
	TokenType            string `json:"token_type,omitempty"`
	ExpiresIn            uint64 `json:"expires_in,omitempty"`
	RefreshToken         string `json:"refresh_token,omitempty"`
	Scope                string `json:"scope,omitempty"`
	IdToken              string `json:"id_token,omitempty"`
	RefreshTokenExpireIn uint64 `json:"refresh_token_expire_in,omitempty"`
}

type socialBase struct {
	Platform string
	Builder  *sql.SCore
	ICore    iCore
}

func load(platform string) iCore {
	key := strings.ToUpper(platform)
	vars.AppId, _ = os.LookupEnv(fmt.Sprintf("%s_APP_ID", key))
	vars.ClientId, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_ID", key))
	appUrl, _ := os.LookupEnv("APP_URL")
	vars.RedirectUri = appUrl + "/api/handle/auth"
	vars.ClientSecret, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_SECRET", key))
	vars.Tenant, _ = os.LookupEnv(fmt.Sprintf("%s_TENANT", key))
	vars.EndPoint, _ = os.LookupEnv(fmt.Sprintf("%s_BASE_API", key))
	vars.Version, _ = os.LookupEnv(fmt.Sprintf("%s_VERSION", key))

	var typeStruct iCore
	switch platform {
	case google:
		typeStruct = sGoogle{}
		break
	case facebook:
		typeStruct = sFacebook{}
		break
	case instagram:
		typeStruct = sInstagram{}
		break
	case github:
		typeStruct = sGithub{}
		break
	case twitter:
		typeStruct = sTwitter{}
		break
	case bitbucket:
		typeStruct = sBitbucket{}
		break
	case dropbox:
		typeStruct = sDropbox{}
		break
	case gitlab:
		typeStruct = sGitlab{}
		break
	case line:
		typeStruct = sLine{}
		break
	case linkedin:
		typeStruct = sLinkedin{}
		break
	case microsoft:
		typeStruct = sMicrosoft{}
		break
	case pinterest:
		typeStruct = sPinterest{}
		break
	case reddit:
		typeStruct = sReddit{}
		break
	case shopify:
		typeStruct = sShopify{}
		break
	case tiktok:
		typeStruct = sTiktok{}
		break
	default:
		helpers.CheckNilErr(errors.New("Platform not found"))
	}
	return typeStruct
}

func New() social {
	platform := vars.Payload.Platform
	social := load(platform)
	social.loadConfig()
	codeVerifier = getCodeVerifier()
	return &socialBase{
		Platform: strings.ToLower(platform),
		Builder:  new(sql.SCore),
		ICore:    social,
	}
}

func (s *socialBase) Generate() string {
	if coreConfig.UsePKCE {
		vars.Payload.CodeVerifier = codeVerifier
	}
	return s.buildLinkAuth(helpers.EncodeJWT(vars.Payload, false))
}
func (s *socialBase) Auth(r *http.Request) {
	code := r.URL.Query().Get("code")
	if coreConfig.UsePKCE {
		code = code + "," + r.Header.Get("code_verifier")
	}
	token := s.ICore.getToken(code)
	if !token.Status {
		helpers.CheckNilErr(token.Error)
		return
	}
	var parseToken resToken
	_ = json.Unmarshal(token.Data, &parseToken)
	idToken = parseToken.IdToken
	coreModel := s.ICore.profile(parseToken.AccessToken)
	if coreModel.InternalId == "" {
		helpers.CheckNilErr(errors.New("authentication failed"))
		return
	}
	coreModel.AccessToken = parseToken.AccessToken
	coreModel.RefreshToken = parseToken.RefreshToken
	coreModel.ExpireToken = time.Now().Add(time.Duration(parseToken.ExpiresIn) * time.Millisecond)
	coreModel.ID = uuid.New()
	s.Builder.ModelBase = &coreModel
	result := s.Builder.UpdateOrCreate()
	if !result.Status {
		helpers.CheckNilErr(result.Errors)
		return
	}
	helpers.FilterDataPrivate(&coreModel)
	b, _ := json.Marshal(coreModel)
	log.Println(string(b))
	config.Pusher(string(b), r.Header.Get("ip"))
}
func (s *socialBase) buildLinkAuth(state string) string {
	queryData := url.Values{}
	queryData.Add("client_id", vars.ClientId)
	queryData.Add("redirect_uri", vars.RedirectUri)
	queryData.Add("scope", s.formatScope())
	queryData.Add("response_type", "code")
	queryData.Add("state", state)
	if coreConfig.UsePKCE {
		queryData.Add("code_challenge", s.getCodeChallenge())
		queryData.Add("code_challenge_method", "S256")
	}
	for k, v := range parameters {
		queryData.Add(k, v)
	}
	return s.getUrlAuth() + "?" + queryData.Encode()
}
func (s *socialBase) getUrlAuth() string {
	return urlAuth
}

func (s *socialBase) formatScope() string {
	log.Println(coreConfig.Scopes)
	return strings.Join(coreConfig.Scopes, coreConfig.Separator)
}
func getCodeVerifier() string {
	buf, _ := helpers.RandomBytes(32)
	return helpers.EncodeS256(buf)
}

func (s *socialBase) getCodeChallenge() string {
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	return helpers.EncodeS256(h.Sum(nil))
}

func buildPayloadToken(code string, typeValue bool) (url.Values, map[string]interface{}) {
	if typeValue {
		queryData := url.Values{}
		queryData.Add("code", code)
		queryData.Add("grant_type", "authorization_code")
		queryData.Add("redirect_uri", vars.RedirectUri)
		queryData.Add("client_id", vars.ClientId)
		queryData.Add("client_secret", vars.ClientSecret)
		if coreConfig.UsePKCE {
			result := strings.Split(code, ",")
			queryData.Del("code")
			queryData.Add("code", result[0])
			queryData.Add("code_verifier", result[1])
		}
		return queryData, nil
	}
	data := make(map[string]interface{})
	data["code"] = code
	data["grant_type"] = "authorization_code"
	data["redirect_uri"] = vars.RedirectUri
	data["client_id"] = vars.ClientId
	data["client_secret"] = vars.ClientSecret
	if coreConfig.UsePKCE {
		result := strings.Split(code, ",")
		data["code"] = result[0]
		data["code_verifier"] = result[1]
	}
	return nil, data

}

func buildPayloadRefresh(refresh string) map[string]string {
	var data map[string]string
	data["grant_type"] = "refresh_token"
	data["refresh_token"] = refresh
	data["redirect_uri"] = vars.RedirectUri
	data["client_id"] = vars.ClientId
	data["client_secret"] = vars.ClientSecret
	return data
}

func headerAuthBasic() map[string]string {
	var headers = make(map[string]string)
	headers["Authorization"] = "Basic " + helpers.Base64Encode(vars.ClientId+":"+vars.ClientSecret)
	return headers
}
