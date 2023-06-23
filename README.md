<p align="center" style="width: 200px;display: flex;align-items: center;justify-content: center"><a href="#" target="_blank"><img src="https://i.ibb.co/vmC7sHn/pngwing-com.png" width="400"></a></p>
<p style="align-items: center; margin:5px auto;display: flex;justify-content: center">Social authentication multiple-platform and management users</p>

## feature
- Facebook
- Google (future)
- Tiktok(future)
- Instagram(future)
- Twitter(future)
- Github(future)
- Linkedin(future)
- Bitbucket(future)
- GitLab(future)
- Microsoft(future)
- Dropbox(future)
- Reddit(future)
- Pinterest(future)
- Line(future)
- shopify(future)
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
go get https://github.com/tranvannghia021/gocore
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
DB_CORE_HOST=
DB_CORE_PORT=5432
DB_CORE_USER=default
DB_CORE_PASSWORD=secret
DB_CORE_NAME=

CORE_REDIS_HOST=
CORE_REDIS_PASSWORD=
CORE_REDIS_PORT=6379
CORE_REDIS_DB=

FACEBOOK_CLIENT_ID=
FACEBOOK_CLIENT_SECRET=
FACEBOOK_BASE_API=https://graph.facebook.com
FACEBOOK_VERSION=v16.0

KEY_JWT=
TIME_EXPIRE=

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
    authRouter.Use(middlewaresCore.VerifyState)
    authRouter.HandleFunc("/auth", gocore.Auth).Methods("GET")
}

```
## API :

| Method  | URI |  Action | Middleware |
|---------| --- | --- | --- |
| POST    | api/{platform}/generate-url |  gocore.GenerateUrl |   |
| GET     | api/handle/auth |  gocore.Auth | middlewaresCore.VerifyState
