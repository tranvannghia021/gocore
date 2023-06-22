package src

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	core "github.com/tranvannghia021/gocore"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var payload = src.PayloadGenerate{}
var codeVerifier string
var urlAuth string

var coreConfig src.ConfigSocial

func New(platform string) string {
	initConfig := func() string {
		return "new" + strings.Title(platform)
	}
	platform = strings.ToUpper(platform)
	core.AppId, _ = os.LookupEnv(fmt.Sprintf("%s_APP_ID", platform))
	core.ClientId, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_ID", platform))
	core.RedirectUri, _ = os.LookupEnv(fmt.Sprintf("%s_REDIRECT_URI", platform))
	core.ClientSecret, _ = os.LookupEnv(fmt.Sprintf("%s_CLIENT_SECRET", platform))
	core.EndPoint, _ = os.LookupEnv(fmt.Sprintf("%s_BASE_API", platform))
	core.Version, _ = os.LookupEnv(fmt.Sprintf("%s_VERSION", platform))
	codeVerifier = getCodeVerifier()
	initConfig()
	return generate()
}

func generate() string {
	if coreConfig.UsePKCE {
		payload.CodeVerifier = codeVerifier
	}
	return buildLinkAuth(helpers.EncodeJWT(payload))
}
func buildLinkAuth(state string) string {
	queryData := url.Values{}
	queryData.Add("client_id", core.ClientId)
	queryData.Add("redirect_uri", core.RedirectUri)
	queryData.Add("scope", formatScope())
	queryData.Add("response_type", "code")
	queryData.Add("state", state)
	queryData.Add("code_challenge", getCodeChallenge())
	queryData.Add("code_challenge_method", "S256")

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
