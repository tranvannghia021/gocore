package socials

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tranvannghia021/gocore/config"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/repositories/sql"
	"github.com/tranvannghia021/gocore/vars"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Social interface {
	Generate() string
	Auth(code string, r *http.Request)
}

var codeVerifier string
var urlAuth string
var parameters = make(map[string]string)
var coreConfig repositories.ConfigSocial

type ResToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	TokenType    string `json:"token_type,omitempty"`
	ExpiresIn    uint64 `json:"expires_in,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type ResProfile struct {
	Id       string
	Name     string
	FistName string
	LastName string
	Email    string
	Picture  struct {
	}
}

type SocialBase struct {
	Platform string
	Builder  *sql.SCore
}

func New(platform string) *SocialBase {
	handler, ok := vars.CallConfig[platform]
	if !ok {
		panic("platform not found")
	}
	platform = strings.ToUpper(platform)
	vars.AppId, _ = os.LookupEnv(fmt.Sprintf("%s_APP_ID", platform))
	vars.ClientId, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_ID", platform))
	appUrl, _ := os.LookupEnv("APP_URL")
	vars.RedirectUri = appUrl + "/api/handle/auth"
	vars.ClientSecret, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_SECRET", platform))
	vars.Tenant, _ = os.LookupEnv(fmt.Sprintf("%s_TENANT", platform))
	vars.EndPoint, _ = os.LookupEnv(fmt.Sprintf("%s_BASE_API", platform))
	vars.Version, _ = os.LookupEnv(fmt.Sprintf("%s_VERSION", platform))
	handler()
	codeVerifier = getCodeVerifier()
	return &SocialBase{
		Platform: strings.ToLower(platform),
		Builder:  &sql.SCore{},
	}
}

func (s *SocialBase) Generate() string {
	if coreConfig.UsePKCE {
		vars.Payload.CodeVerifier = codeVerifier
	}
	return buildLinkAuth(helpers.EncodeJWT(vars.Payload, false))
}
func (s *SocialBase) Auth(code string, r *http.Request) {
	getToken, _ := vars.PLatFormToken[s.Platform]
	token := getToken(code)
	if !token.Status {
		helpers.CheckNilErr(token.Error)
		return
	}
	var parseToken ResToken
	_ = json.Unmarshal(token.Data, &parseToken)
	getProfile, _ := vars.PLatFormProfile[s.Platform]
	coreModel := getProfile(parseToken.AccessToken)
	coreModel.AccessToken = parseToken.AccessToken
	coreModel.RefreshToken = parseToken.RefreshToken
	coreModel.ExpireToken = time.Now().Add(time.Duration(parseToken.ExpiresIn) * time.Millisecond)
	coreModel.ID = uuid.New()
	result := s.Builder.UpdateOrCreate(&coreModel)
	if !result.Status {
		helpers.CheckNilErr(result.Errors)
		return
	}
	helpers.FilterDataPrivate(&coreModel)
	b, _ := json.Marshal(coreModel)
	config.Pusher(string(b), r.Header.Get("ip"))
}
func buildLinkAuth(state string) string {
	queryData := url.Values{}
	queryData.Add("client_id", vars.ClientId)
	queryData.Add("redirect_uri", vars.RedirectUri)
	queryData.Add("scope", formatScope())
	queryData.Add("response_type", "code")
	queryData.Add("state", state)
	if coreConfig.UsePKCE {
		queryData.Add("code_challenge", getCodeChallenge())
		queryData.Add("code_challenge_method", "S256")
	}
	for k, v := range parameters {
		queryData.Add(k, v)
	}
	return getUrlAuth() + "?" + queryData.Encode()
}
func getUrlAuth() string {
	return urlAuth
}

func formatScope() string {
	return strings.Join(coreConfig.Scopes, coreConfig.Separator)
}
func getCodeVerifier() string {
	buf := make([]byte, 128)
	_, err := rand.Read(buf)
	helpers.CheckNilErr(err)
	val, _ := bin2hex(string(buf))
	return val
}

func getCodeChallenge() string {
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	bs := h.Sum(nil)
	replace := make(map[string]string)
	replace["+/"] = "-_"
	return strings.Trim(strtr(base64.StdEncoding.EncodeToString(bs), replace), "=")
}

func bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

func strtr(str string, replace map[string]string) string {
	if len(replace) == 0 || len(str) == 0 {
		return str
	}
	for old, new := range replace {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
}

func buildPayloadToken(code string) string {
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
	return queryData.Encode()
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
