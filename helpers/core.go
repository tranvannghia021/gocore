package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tranvannghia021/gocore/src/repositories"
	"os"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

var keyJwt, _ = os.LookupEnv("KEY_JWT")
var jwtKey = []byte(keyJwt)

func EncodeJWT(payload repositories.PayloadGenerate) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	payload.CreateAt = time.Now().Add(10 * time.Minute)
	claims["ID"] = payload.ID
	claims["Email"] = payload.Email
	claims["CreateAt"] = payload.CreateAt
	claims["CreateAt"] = payload.CreateAt
	claims["Ip"] = payload.Ip
	claims["CodeVerifier"] = payload.CodeVerifier
	claims["Platform"] = payload.Platform

	tokenString, err := token.SignedString(jwtKey)
	CheckNilErr(err)
	return tokenString
}
func IsExpire(timeState time.Time) bool {
	return timeState.Before(time.Now())
}

func DecodeJWT(tokenString string) (repositories.PayloadGenerate, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(keyJwt), nil
	})
	claims, ok := token.Claims.(*jwt.MapClaims)
	if ok && token.Valid {

	} else {
		panic(err)
	}
	jsonString, _ := json.Marshal(claims)
	fmt.Println(string(jsonString))
	var payload repositories.PayloadGenerate
	_ = json.Unmarshal(jsonString, &payload)

	return payload, IsExpire(payload.CreateAt)
}
