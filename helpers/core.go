package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tranvannghia021/gocore/src/repositories"
	"github.com/tranvannghia021/gocore/vars"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		//var w io.Writer
		panic(err)
		//var w http.ResponseWriter
		//fmt.Fprintf(w, "Invalid request body error:%sâ€", err.Error())
		//json.NewEncoder(w).Encode("adsaas")
	}
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
		panic(err)
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
