package src

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tranvannghia021/gocore/helpers"
	"github.com/tranvannghia021/gocore/src/mail"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/src/repositories/sql"
	"github.com/tranvannghia021/gocore/src/response"
	"github.com/tranvannghia021/gocore/src/socials"
	"github.com/tranvannghia021/gocore/vars"
	"log"
	"net/http"
	"os"
	"time"
)

var inApp = "APP"

type iHandler interface {
	GenerateUrl(w http.ResponseWriter, r *http.Request)
	AuthHandle(w http.ResponseWriter, r *http.Request)
}
type shandler struct {
}
type styleRpPusher struct {
	Url    string `json:"url"`
	Pusher struct {
		Channel string `json:"channel"`
		Event   string `json:"event"`
	} `json:"pusher"`
}
type sShopify struct {
	Domain string `json:"domain"`
}

func NewHandler() iHandler {
	return &shandler{}
}
func (shandler) GenerateUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vars.Payload.Ip = helpers.Md5(r.RemoteAddr)
	vars.Payload.Platform = params["platform"]
	var s sShopify
	_ = json.NewDecoder(r.Body).Decode(&s)
	vars.Payload.Domain = s.Domain
	channel, _ := os.LookupEnv("CHANNEL_NAME")
	event, _ := os.LookupEnv("EVENT_NAME")
	response.Response(&styleRpPusher{
		Url: socials.New().Generate(),
		Pusher: struct {
			Channel string `json:"channel"`
			Event   string `json:"event"`
		}{
			Channel: channel,
			Event:   event + "_" + vars.Payload.Ip,
		},
	}, w, false, nil)
}

func (shandler) AuthHandle(w http.ResponseWriter, r *http.Request) {
	vars.Payload.Platform = r.Header.Get("platform")
	params := r.URL.Query()
	vars.Payload.Domain = params.Get("shop")
	if params.Get("error") != "" || params.Get("errors") != "" {
		helpers.CheckNilErr(errors.New("Accept denied!"))
		return
	}
	socials.New().Auth(r)
	fontEndUrl, _ := os.LookupEnv("FONT_END_URL")
	http.Redirect(w, r, fontEndUrl, 302)
}

type iManager interface {
	List(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Me(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
	Resend(w http.ResponseWriter, r *http.Request)
	Verify(w http.ResponseWriter, r *http.Request)
	Success(w http.ResponseWriter, r *http.Request)
}

type sManager struct {
}

func NewManager() iManager {
	return &sManager{}
}

type payloadBody struct {
	Email     string    `json:"email,omitempty" validate:"required,email"`
	Password  string    `json:"password,omitempty" validate:"required,min=8,max=32"`
	Confirm   string    `json:"confirm,omitempty" validate:"required,min=8,max=32"`
	FirstName string    `json:"first_name,omitempty" validate:"required"`
	LastName  string    `json:"last_name,omitempty"  validate:"required"`
	Avatar    string    `json:"avatar,omitempty"`
	Gender    string    `json:"gender,omitempty"  validate:"required"`
	Phone     string    `json:"phone,omitempty"  validate:"required"`
	BirthDay  time.Time `json:"birth_day" `
	Address   string    `json:"address,omitempty"`
}

func (sManager) List(w http.ResponseWriter, r *http.Request) {

}

func (sManager) SignIn(w http.ResponseWriter, r *http.Request) {
	var core = repositories.Core{}
	_ = json.NewDecoder(r.Body).Decode(&core)
	password := core.Password
	core.Platform = inApp
	sql := new(sql.Suser)
	sql.ModelBase = &core
	record := sql.First()
	if !record.Status {
		response.Response(nil, w, true, record.Errors)
		return
	}
	if !core.Status {
		response.Response(nil, w, true, errors.New("Please verify your email"))
		return
	}
	if !helpers.CheckPasswordHash(password, core.Password) {
		response.Response(nil, w, true, errors.New("Password does not match"))
		return
	}
	response.Response(helpers.BuildResPayloadJwt(core, true), w, false, nil)
}

func (sManager) SignUp(w http.ResponseWriter, r *http.Request) {
	var body payloadBody
	_ = json.NewDecoder(r.Body).Decode(&body)
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		response.Response(nil, w, true, err)
		return
	}
	if body.Password != body.Confirm {
		response.Response(nil, w, true, errors.New("Password does not match"))
		return
	}
	var pathAvatar string
	if body.Avatar != "" {
		path, err := helpers.Base64ToImage(body.Avatar, "./avatar/")
		helpers.CheckNilErr(err)
		pathAvatar = path
	}
	uuidA := uuid.New()
	core := repositories.Core{
		ID:            uuidA,
		InternalId:    uuidA.String(),
		Platform:      inApp,
		Email:         body.Email,
		EmailVerifyAt: time.Time{},
		Password:      helpers.HashPassword(body.Password),
		FirstName:     body.FirstName,
		LastName:      body.LastName,
		Avatar:        pathAvatar,
		Gender:        body.Gender,
		Status:        false,
		Phone:         body.Phone,
		BirthDay:      body.BirthDay,
		Address:       body.Address,
	}
	var repo sql.Suser
	repo.ModelBase = &core
	result := repo.Create()
	if !result.Status {
		response.Response(nil, w, true, result.Errors)
		return
	}
	var payload = vars.PayloadGenerate{
		Platform: inApp,
		ID:       core.ID,
		Email:    core.Email,
		CreateAt: time.Now().Add(10 * time.Minute),
	}
	signature := helpers.EncodeJWT(payload, false)
	appurl, _ := os.LookupEnv("APP_URL")
	go log.Println("MAIL", mail.RegisterMail(core.Email, appurl+"/api/verify/email?token="+signature))
	helpers.FilterDataPrivate(&core)
	response.Response(core, w, false, nil)
}
func (sManager) Update(w http.ResponseWriter, r *http.Request) {

}
func (sManager) Me(w http.ResponseWriter, r *http.Request) {
	var core = repositories.Core{ID: uuid.MustParse(r.Header.Get("id"))}
	sql := new(sql.Suser)
	sql.ModelBase = &core
	record := sql.First()
	if !record.Status {
		response.Response(nil, w, true, record.Errors)
		return
	}
	helpers.FilterDataPrivate(&core)
	response.Response(core, w, false, nil)
}
func (sManager) Delete(w http.ResponseWriter, r *http.Request) {
	var core = repositories.Core{ID: uuid.MustParse(r.Header.Get("id"))}
	sql := new(sql.Suser)
	sql.ModelBase = &core
	record := sql.Delete()
	if !record.Status {
		response.Response(nil, w, true, record.Errors)
		return
	}
	response.Response(nil, w, false, nil)
}

func (sManager) Refresh(w http.ResponseWriter, r *http.Request) {
	var core = repositories.Core{ID: uuid.MustParse(r.Header.Get("id"))}
	sql := new(sql.Suser)
	sql.ModelBase = &core
	record := sql.First()
	if !record.Status {
		response.Response(nil, w, true, record.Errors)
		return
	}
	response.Response(helpers.BuildResPayloadJwt(core, false), w, false, nil)
}

func (sManager) Resend(w http.ResponseWriter, r *http.Request) {

}

func (sManager) Verify(w http.ResponseWriter, r *http.Request) {

	uuid, _ := uuid.FromBytes([]byte(r.Header.Get("id")))
	userModel := new(sql.Suser)
	core := repositories.Core{
		ID:       uuid,
		Platform: inApp,
		Email:    r.Header.Get("email"),
	}
	userModel.ModelBase = &core
	record := userModel.First()
	if !record.Status {
		response.Response(nil, w, true, record.Errors)
		return
	}
	core.Status = true
	core.IsDisconnect = false
	userModel.Update()
	http.Redirect(w, r, "success", 302)
}

func (sManager) Success(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/"))
}
