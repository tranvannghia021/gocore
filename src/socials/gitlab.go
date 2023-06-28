package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"net/url"
	"strconv"
	"time"
)

var gitlab = "gitlab"

type sGitlab struct {
}

var scopeGl []string

type profileGl struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	Name            string    `json:"name"`
	State           string    `json:"state"`
	AvatarURL       string    `json:"avatar_url"`
	WebURL          string    `json:"web_url"`
	CreatedAt       time.Time `json:"created_at"`
	Bio             string    `json:"bio"`
	Location        string    `json:"location"`
	PublicEmail     any       `json:"public_email"`
	Skype           string    `json:"skype"`
	Linkedin        string    `json:"linkedin"`
	Twitter         string    `json:"twitter"`
	Discord         string    `json:"discord"`
	WebsiteURL      string    `json:"website_url"`
	Organization    string    `json:"organization"`
	JobTitle        string    `json:"job_title"`
	Pronouns        any       `json:"pronouns"`
	Bot             bool      `json:"bot"`
	WorkInformation any       `json:"work_information"`
	LocalTime       any       `json:"local_time"`
	LastSignInAt    time.Time `json:"last_sign_in_at"`
	ConfirmedAt     time.Time `json:"confirmed_at"`
	LastActivityOn  string    `json:"last_activity_on"`
	Email           string    `json:"email"`
	ThemeID         int       `json:"theme_id"`
	ColorSchemeID   int       `json:"color_scheme_id"`
	ProjectsLimit   int       `json:"projects_limit"`
	CurrentSignInAt time.Time `json:"current_sign_in_at"`
	Identities      []struct {
		Provider       string `json:"provider"`
		ExternUID      string `json:"extern_uid"`
		SamlProviderID any    `json:"saml_provider_id"`
	} `json:"identities"`
	CanCreateGroup                 bool   `json:"can_create_group"`
	CanCreateProject               bool   `json:"can_create_project"`
	TwoFactorEnabled               bool   `json:"two_factor_enabled"`
	External                       bool   `json:"external"`
	PrivateProfile                 bool   `json:"private_profile"`
	CommitEmail                    string `json:"commit_email"`
	SharedRunnersMinutesLimit      any    `json:"shared_runners_minutes_limit"`
	ExtraSharedRunnersMinutesLimit any    `json:"extra_shared_runners_minutes_limit"`
	ScimIdentities                 []any  `json:"scim_identities"`
}

func (s sGitlab) loadConfig() {
	coreConfig.Separator = " "
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"read_user",
	}, scopeGl...))
	urlAuth = fmt.Sprintf("%s/oauth/authorize", vars.EndPoint)
}

func (s sGitlab) getToken(code string) vars.ResReq {
	body, _ := buildPayloadToken(code, true)
	return service.PostRequest(fmt.Sprintf("%s/oauth/token?%s", vars.EndPoint, body.Encode()), nil, nil)
}

func (s sGitlab) profile(token string) repositories.Core {
	query := url.Values{}
	query.Add("access_token", token)
	result := service.GetRequest(fmt.Sprintf("%s/api/%s/user?%s", vars.EndPoint, vars.Version, query.Encode()), nil)
	if !result.Status {
		helpers.CheckNilErr(result.Error)
		return repositories.Core{}
	}
	var profile profileGl
	_ = json.Unmarshal(result.Data, &profile)
	return repositories.Core{
		InternalId:    strconv.Itoa(profile.ID),
		Platform:      gitlab,
		Email:         profile.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name,
		Avatar:        profile.AvatarURL,
		Status:        true,
	}
}

func AddScopeGitlab(scope []string) {
	scopeGl = scope
}
