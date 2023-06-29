<p align="center" style="width: 200px;display: flex;align-items: center;justify-content: center"><a href="#" target="_blank"><img src="https://i.ibb.co/vmC7sHn/pngwing-com.png" width="400"></a></p>
<p style="align-items: center; margin:5px auto;display: flex;justify-content: center">Social authentication multiple-platform and management users</p>

## feature
- Facebook
- Google
- Instagram
- Twitter
- Github
- Linkedin
- Bitbucket
- GitLab
- Microsoft
- Dropbox
- Reddit
- Pinterest(future)
- Line
- shopify
## Official Core SDKs
<div>
<ul>
    <li><a href="https://github.com/tranvannghia021/core">Php</a></li>
    <li><a href="https://github.com/tranvannghia021/gocore">GO</a></li>
</ul>
</div>


## Required
- go >= 1.19
- gorilla/mux >= v1.8
- gorm.io/gorm >= v1.25
## Install
```bash
go get -u github.com/tranvannghia021/gocore
```
## Setup
-    If You want custom scope or field for platform.Ex:

```go
// require config in func Name (init) 
func init() {
	....
		// add sopes in fb
	socials.AddScopeFaceBook([]string{
		"public_profile",
		"email",
		// comtom
	})
    // add field get profile (optional)
	socials.AddFieldFacebook([]string{
		"id",
		"name",
		"first_name",
		"last_name",
		"email",
		"birthday",
		"gender",
		"hometown",
		"location",
		"picture",
		// customs
	})
	....
}

```


- ENV

```env
APP_URL=
FONT_END_URL=

DB_CORE_HOST=
DB_CORE_PORT=
DB_CORE_USER=
DB_CORE_PASSWORD=
DB_CORE_NAME=

CORE_REDIS_HOST=
CORE_REDIS_PASSWORD=
CORE_REDIS_PORT=
CORE_REDIS_DB=

FACEBOOK_CLIENT_ID=
FACEBOOK_CLIENT_SECRET=
FACEBOOK_BASE_API=
FACEBOOK_VERSION=

GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_BASE_API=
GOOGLE_VERSION=

GITHUB_APP_ID=
GITHUB_CLIENT_ID=
GITHUB_CLIENT_SECRET=
GITHUB_BASE_API=https://api.github.com
GITHUB_VERSION=v1

TIKTOK_APP_ID=
TIKTOK_CLIENT_ID=
TIKTOK_CLIENT_SECRET=
TIKTOK_BASE_API=https://open.tiktokapis.com
TIKTOK_VERSION=v2

TWITTER_CLIENT_ID=
TWITTER_CLIENT_SECRET=
TWITTER_BASE_API=https://api.twitter.com
TWITTER_VERSION=2

INSTAGRAM_CLIENT_ID=
INSTAGRAM_CLIENT_SECRET=
INSTAGRAM_BASE_API=https://graph.instagram.com

LINKEDIN_CLIENT_ID=
LINKEDIN_CLIENT_SECRET=
LINKEDIN_BASE_API=https://api.linkedin.com
LINKEDIN_VERSION=v2

BITBUCKET_CLIENT_ID=
BITBUCKET_CLIENT_SECRET=
BITBUCKET_BASE_API=https://api.bitbucket.org
BITBUCKET_VERSION=2.0

GITLAB_APP_ID=
GITLAB_CLIENT_ID=${GITLAB_APP_ID}
GITLAB_CLIENT_SECRET=
GITLAB_BASE_API=https://gitlab.com
GITLAB_VERSION=v4

MICROSOFT_APP_ID=
MICROSOFT_TENANT=
MICROSOFT_CLIENT_ID=${MICROSOFT_APP_ID}
MICROSOFT_CLIENT_SECRET=
MICROSOFT_BASE_API=https://graph.microsoft.com
MICROSOFT_VERSION=v1.0

DROPBOX_CLIENT_ID=
DROPBOX_CLIENT_SECRET=
DROPBOX_BASE_API=https://api.dropboxapi.com
DROPBOX_VERSION=2

REDDIT_CLIENT_ID=
REDDIT_CLIENT_SECRET=
REDDIT_BASE_API=https://oauth.reddit.com
REDDIT_VERSION=v1

PINTEREST_APP_ID=
PINTEREST_CLIENT_ID=${PINTEREST_APP_ID}
PINTEREST_CLIENT_SECRET=
PINTEREST_BASE_API=https://api.pinterest.com
PINTEREST_VERSION=v5

LINE_APP_ID=
LINE_CLIENT_ID=${LINE_APP_ID}
LINE_CLIENT_SECRET=
LINE_BASE_API=https://api.line.me
LINE_VERSION=v2.1

SHOPIFY_APP_ID=
SHOPIFY_CLIENT_ID=
SHOPIFY_CLIENT_SECRET=
SHOPIFY_VERSION=2023-01


KEY_JWT=
TIME_EXPIRE=
KEY_PRIVATE_JWT=
TIME_PRIVATE_EXPIRE=

CHANNEL_NAME=
EVENT_NAME=
PUSHER_APP_ID=
PUSHER_APP_KEY=
PUSHER_APP_SECRET=
PUSHER_APP_CLUSTER=ap1


```
- In the mux.router
```go
func apiRouter(router *mux.Router) {
    router.HandleFunc("/{platform}/generate-url", gocore.GenerateUrl).Methods("POST")
    authRouter := router.PathPrefix("/handle").Subrouter()
    authRouter.Use(middlewares.VerifyState)
    authRouter.Use(middlewares.VerifyHmac) //shopify
    authRouter.HandleFunc("/auth", gocore.Auth).Methods("GET")
}

```
## API :

| Method  | URI                         | Action             | Middleware                                       |
|---------|-----------------------------|--------------------|--------------------------------------------------|
| POST    | api/{platform}/generate-url | gocore.GenerateUrl |                                                  |
| GET     | api/handle/auth             | gocore.Auth        | middlewares.VerifyState ,middlewares.VerifyHmac  |


## Config Scope
- Structure

| -      | Scope                                    | fields                                    |
|--------|------------------------------------------|-------------------------------------------|
|        | AddScope{platform}(scope []string)       | AddField{platform}(fields []string)       |
| VD     | socials.AddScopeFacebook(scope []string) | socials.AddFieldFacebook(fields []string) |

- Support

| platform  | scopes | fields(profile) |
|-----------|--------|-----------------|
| Facebook  | Yes    | Yes             |    
| Google    | Yes    | No              |
| Bitbucket | Yes    | No              |
| Dropbox   | Yes    | No              |
| Github    | Yes    | No              |
| Gitlab    | Yes    | No              |
| Line      | Yes    | No              |
| Linkedin  | Yes    | No              |
| Microsoft | Yes    | No              |
| Reddit    | Yes    | No              |
| Twitter   | Yes    | Yes             |
| Instagram | Yes    | Yes             |
