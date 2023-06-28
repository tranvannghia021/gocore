package socials

var tiktok = "tiktok"

type sTiktok struct {
}

//
//var tiktok string = "tiktok"
//var defaultScopeTt = []string{
//	"user.info.basic",
//}
//var scopeTt []string
//var defaultFieldsTt = []string{
//	"open_id",
//	"union_id",
//	"avatar_url",
//	"display_name",
//	"bio_description",
//	"profile_deep_link",
//	"is_verified",
//	"follower_count",
//	"following_count",
//	"likes_count",
//	"video_count",
//}
//var fieldTt []string
//var loadConfigTt = func() {
//	coreConfig.UsePKCE = false
//	coreConfig.Separator = ","
//	urlAuth = "https://www.tiktok.com/v2/auth/authorize/"
//}
//
//func init() {
//	AddScopeTiktok(scopeTt)
//	AddFieldTiktok(fieldTt)
//	vars.CallConfig[tiktok] = loadConfigTt
//	vars.PLatFormToken[tiktok] = getTokenTt
//	vars.PLatFormProfile[tiktok] = profileTt
//}
//
//func AddScopeTiktok(scope []string) {
//	scopeTt = scope
//	coreConfig.Scopes = helpers.RemoveDuplicateStr(append(defaultScopeTt, scope...))
//}
//
//func AddFieldTiktok(fields []string) {
//	fieldTt = fields
//	coreConfig.Fields = helpers.RemoveDuplicateStr(append(defaultFieldsTt, fields...))
//}
//
//var getTokenTt = func(code string) vars.ResReq {
//	_, data := buildPayloadToken(code, false)
//	url := fmt.Sprintf("%s/%s/oauth/token/", vars.EndPoint, vars.Version)
//	var headers = make(map[string]string)
//	headers["Content-Type"] = "application/x-www-form-urlencoded"
//	headers["Cache-Control"] = "no-cache"
//	return service.PostRequest(url, headers, data)
//}
//
//var profileTt = func(token string) repositories.Core {
//
//	results := service.GetRequest(fmt.Sprintf("%s/%s/user/info/%s",
//		vars.EndPoint,
//		vars.Version, strings.Join(fieldTt, ",")), nil)
//	if !results.Status {
//		helpers.CheckNilErr(results.Error)
//	}
//	var profile ProfileFB
//	_ = json.Unmarshal(results.Data, &profile)
//	return repositories.Core{
//		InternalId:    profile.ID,
//		Platform:      facebook,
//		Email:         profile.Email,
//		EmailVerifyAt: time.Now(),
//		Password:      "",
//		FirstName:     profile.FirstName,
//		LastName:      profile.LastName,
//		Avatar:        profile.Picture.Data.URL,
//		Status:        true,
//	}
//}
//
//type ProfileTt struct {
//	ID        string `json:"id"`
//	Name      string `json:"name"`
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//	Email     string `json:"email"`
//	Picture   struct {
//		Data struct {
//			Height       int    `json:"height"`
//			IsSilhouette bool   `json:"is_silhouette"`
//			URL          string `json:"url"`
//			Width        int    `json:"width"`
//		} `json:"data"`
//	} `json:"picture"`
//}
