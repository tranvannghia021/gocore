package helpers

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tranvannghia021/gocore/src/repositories"
	"os"
	"strconv"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FilterDataPrivate(coreModel *repositories.Core) {
	coreModel.AccessToken = ""
	coreModel.RefreshToken = ""
	coreModel.Password = ""
	coreModel.EmailVerifyAt = time.Time{}
	coreModel.ExpireToken = time.Time{}

}
func EncodeJWT(payload repositories.PayloadGenerate) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	timeJwt, _ := strconv.Atoi(getTimeExpire()) //time.Duration(timeJwt) *
	payload.CreateAt = time.Now().Add(time.Duration(timeJwt) * time.Minute)
	claims["ID"] = payload.ID
	claims["Email"] = payload.Email
	claims["CreateAt"] = payload.CreateAt
	claims["CreateAt"] = payload.CreateAt
	claims["Ip"] = payload.Ip
	claims["CodeVerifier"] = payload.CodeVerifier
	claims["Platform"] = payload.Platform

	tokenString, err := token.SignedString(getKeyJWT())
	CheckNilErr(err)
	return tokenString
}
func isExpire(timeState time.Time) bool {
	return timeState.Before(time.Now())
}

func DecodeJWT(tokenString string) (repositories.PayloadGenerate, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return getKeyJWT(), nil
	})
	claims, _ := token.Claims.(*jwt.MapClaims)
	if !token.Valid {
		panic(err)
	}
	jsonString, _ := json.Marshal(claims)
	var payload repositories.PayloadGenerate
	_ = json.Unmarshal(jsonString, &payload)
	return payload, isExpire(payload.CreateAt)
}

func getKeyJWT() []byte {
	keyJwt, _ := os.LookupEnv("KEY_JWT")
	return []byte(keyJwt)
}

func getTimeExpire() string {
	timeExpire, _ := os.LookupEnv("TIME_EXPIRE")
	return timeExpire
}
