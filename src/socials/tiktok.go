package socials

import (
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"log"
	"strings"
)

var tiktok = "tiktok"

type sTiktok struct {
}

var (
	scopeTiktok  []string
	fieldsTiktok []string
)

func (s sTiktok) loadConfig() {
	coreConfig.Separator = ","
	urlAuth = "https://www.tiktok.com/v2/auth/authorize/"
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"user.info.basic",
	}, scopeTiktok...))
	coreConfig.Fields = helpers.RemoveDuplicateStr(append([]string{
		"open_id",
		"union_id",
		"avatar_url",
		"display_name",
		"bio_description",
		"profile_deep_link",
		"is_verified",
		"follower_count",
		"following_count",
		"likes_count",
		"video_count",
	}, fieldsTiktok...))
	parameters["client_key"] = vars.ClientId

}
func (s sTiktok) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	return service.PostFormDataRequest(fmt.Sprintf("%s/%s/oauth/token/", vars.EndPoint, vars.Version), nil, body)
}

func (s sTiktok) profile(token string) repositories.Core {
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	result := service.GetRequest(fmt.Sprintf("%s/%s/user/info/%s", vars.EndPoint, vars.Version, strings.Join(coreConfig.Fields, "&")), headers)
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	log.Println(string(result.Data))
	panic("")
}

func AddScopeTiktok(scope []string) {
	scopeTiktok = scope
}

func AddFieldTiktok(fields []string) {
	fieldsTiktok = fields
}
