package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/vars"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		//var w io.Writer
		//panic(err)
		log.Println("DEBUG_ERR:  ", err)
		//var w http.ResponseWriter
		//fmt.Fprintf(w, "Invalid request body error:%sâ€", err.Error())
		//json.NewEncoder(w).Encode("adsaas")
	}
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
func Strtrim(str string, replace map[string]string) string {
	if len(replace) == 0 || len(str) == 0 {
		return str
	}
	for old, new := range replace {
		str = strings.ReplaceAll(str, old, new)
	}
	return str
}

func EncodeS256(msg []byte) string {
	encoded := base64.StdEncoding.EncodeToString(msg)
	encoded = strings.Replace(encoded, "+", "-", -1)
	encoded = strings.Replace(encoded, "/", "_", -1)
	encoded = strings.Replace(encoded, "=", "", -1)
	return encoded
}

func Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

func FilterDataPrivate(coreModel *repositories.Core) {
	coreModel.AccessToken = ""
	coreModel.RefreshToken = ""
	coreModel.Password = ""
	coreModel.EmailVerifyAt = time.Time{}
	coreModel.ExpireToken = time.Time{}

}
func EncodeJWT(payload vars.PayloadGenerate, isRefresh bool) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var keyByte []byte
	var timeString string
	if isRefresh {
		keyByte = getKeyRefreshJWT()
		timeString = getTimeRefreshExpire()
	} else {
		keyByte = getKeyJWT()
		timeString = getTimeExpire()
	}
	timeJwt, _ := strconv.Atoi(timeString) //time.Duration(timeJwt) *
	payload.CreateAt = time.Now().Add(time.Duration(timeJwt) * time.Minute)
	claims["ID"] = payload.ID
	claims["Email"] = payload.Email
	claims["CreateAt"] = payload.CreateAt
	claims["CreateAt"] = payload.CreateAt
	claims["Ip"] = payload.Ip
	claims["CodeVerifier"] = payload.CodeVerifier
	claims["Platform"] = payload.Platform

	tokenString, err := token.SignedString(keyByte)
	CheckNilErr(err)
	return tokenString
}
func isExpire(timeState time.Time) bool {
	return timeState.Before(time.Now())
}

func DecodeJWT(tokenString string, isRefresh bool) (vars.PayloadGenerate, bool) {
	if strings.Contains(tokenString, "Bearer") {
		tokenString = strings.Split(tokenString, "Bearer ")[1]
	}
	var keyByte []byte
	if isRefresh {
		keyByte = getKeyRefreshJWT()
	} else {
		keyByte = getKeyJWT()
	}
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return keyByte, nil
	})
	claims, _ := token.Claims.(*jwt.MapClaims)
	if !token.Valid {
		log.Fatal(err)
	}
	jsonString, _ := json.Marshal(claims)
	_ = json.Unmarshal(jsonString, &vars.Payload)
	return vars.Payload, isExpire(vars.Payload.CreateAt)
}

func getKeyJWT() []byte {
	keyJwt, _ := os.LookupEnv("KEY_JWT")
	return []byte(keyJwt)
}

func getKeyRefreshJWT() []byte {
	keyJwt, _ := os.LookupEnv("KEY_PRIVATE_JWT")
	return []byte(keyJwt)
}
func getTimeRefreshExpire() string {
	timeExpire, _ := os.LookupEnv("TIME_PRIVATE_EXPIRE")
	return timeExpire
}
func getTimeExpire() string {
	timeExpire, _ := os.LookupEnv("TIME_EXPIRE")
	return timeExpire
}
func CheckTable(db *gorm.DB) bool {

	return db.Migrator().HasTable(&repositories.Core{})
}

func MigrateCore(db *gorm.DB) {
	er := db.Migrator().CreateTable(&repositories.Core{})
	CheckNilErr(er)
}

func RollbackMigrate(db *gorm.DB) {
	er := db.Migrator().DropTable(&repositories.Core{})
	CheckNilErr(er)
}
func Md5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func RandomBytes(length int) ([]byte, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	const csLen = byte(len(charset))
	output := make([]byte, 0, length)
	for {
		buf := make([]byte, length)
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			return nil, fmt.Errorf("failed to read random bytes: %v", err)
		}
		for _, b := range buf {
			// Avoid bias by using a value range that's a multiple of 62
			if b < (csLen * 4) {
				output = append(output, charset[b%csLen])

				if len(output) == length {
					return output, nil
				}
			}
		}
	}

}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkValidType(base64 string) (string, bool) {
	listExt := []string{
		"png",
		"jpeg",
		"jpg",
		"gif",
	}
	for _, v := range listExt {
		if strings.Contains(base64, v) {
			return v, true
		}
	}
	return "", false
}

func Base64ToImage(base64String string, folder string) (string, error) {

	typeImg, ok := checkValidType(base64String)
	if !ok {
		return "", errors.New("Image is invalid")
	}
	b64data := base64String[strings.IndexByte(base64String, ',')+1:]

	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}

	// Create a reader from the byte slice
	reader := strings.NewReader(string(data))

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}
	nameFile := Md5(typeImg+time.Now().String()) + "." + typeImg
	pathFile := folder + nameFile

	file, err := os.Create(pathFile)
	if err != nil {
		return "", err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return "", err
	}

	return nameFile, nil
}

type BuildResLogin struct {
	User repositories.Core `json:"user"`
	Jwt  struct {
		AccessToken   string    `json:"access_token,omitempty"`
		RefreshToken  string    `json:"refresh_token,omitempty"`
		ExpireToken   time.Time `json:"expire_token,omitempty"`
		ExpireRefresh time.Time `json:"expire_refresh,omitempty"`
		Type          string    `json:"type"`
	} `json:"jwt"`
}

func BuildResPayloadJwt(core repositories.Core, isBuildRefresh bool) BuildResLogin {
	timeTokenString, _ := os.LookupEnv("TIME_EXPIRE")
	timeToken, _ := strconv.Atoi(timeTokenString)
	timeRefreshString, _ := os.LookupEnv("TIME_PRIVATE_EXPIRE")
	timeRefresh, _ := strconv.Atoi(timeRefreshString)
	var payload = vars.PayloadGenerate{
		ID:       core.ID,
		Email:    core.Email,
		CreateAt: time.Now().Add(time.Duration(timeToken) * time.Minute),
	}
	payloadRefresh := payload
	payloadRefresh.CreateAt = time.Now().Add(time.Duration(timeRefresh) * time.Minute)
	FilterDataPrivate(&core)
	data := BuildResLogin{
		User: core,
		Jwt: struct {
			AccessToken   string    `json:"access_token,omitempty"`
			RefreshToken  string    `json:"refresh_token,omitempty"`
			ExpireToken   time.Time `json:"expire_token,omitempty"`
			ExpireRefresh time.Time `json:"expire_refresh,omitempty"`
			Type          string    `json:"type"`
		}{
			AccessToken:   EncodeJWT(payload, false),
			RefreshToken:  EncodeJWT(payloadRefresh, true),
			ExpireToken:   payload.CreateAt,
			ExpireRefresh: payloadRefresh.CreateAt,
			Type:          "Bearer",
		},
	}
	if !isBuildRefresh {
		data.Jwt.ExpireRefresh = time.Time{}
		data.Jwt.RefreshToken = ""
	}
	return data
}
