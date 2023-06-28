package socials

import (
	"encoding/json"
	"fmt"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/service"
	"github.com/tranvannghia021/gocore/vars"
	"strconv"
	"time"
)

var github = "github"

type sGithub struct {
}

var scopeGh []string

type profileGh struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Company           any       `json:"company"`
	Blog              string    `json:"blog"`
	Location          any       `json:"location"`
	Email             string    `json:"email"`
	Hireable          any       `json:"hireable"`
	Bio               any       `json:"bio"`
	TwitterUsername   any       `json:"twitter_username"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Plan              struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

func (s sGithub) loadConfig() {
	coreConfig.Separator = ","
	urlAuth = "https://github.com/login/oauth/authorize"
	coreConfig.Scopes = helpers.RemoveDuplicateStr(append([]string{
		"user",
	}, scopeGh...))
}

func (s sGithub) getToken(code string) vars.ResReq {
	data, _ := buildPayloadToken(code, true)
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?%s", data)
	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	return service.PostRequest(url, headers, nil)
}

func (s sGithub) profile(token string) repositories.Core {
	var headers = make(map[string]string)
	headers["Authorization"] = "Bearer " + token
	results := service.GetRequest(fmt.Sprintf("%s/user", vars.EndPoint), headers)
	if !results.Status {
		helpers.CheckNilErr(results.Error)
	}
	var profile profileGh
	_ = json.Unmarshal(results.Data, &profile)
	return repositories.Core{
		InternalId:    strconv.Itoa(profile.ID),
		Platform:      github,
		Email:         profile.Email,
		EmailVerifyAt: time.Now(),
		Password:      "",
		FirstName:     profile.Name,
		Avatar:        profile.AvatarURL,
		Status:        true,
	}
}

func AddScopeGithub(scope []string) {
	scopeGh = scope
}
