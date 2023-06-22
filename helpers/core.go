package helpers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tranvannghia021/gocore/src"
	"time"
)

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

var jwtKey = []byte("ajksdhkajhdkasjbdk")

func EncodeJWT(payload src.PayloadGenerate) string {
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

func DecodeJWT(tokenString string) (src.PayloadGenerate, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ajksdhkajhdkasjbdk"), nil
	})
	claims, ok := token.Claims.(*jwt.MapClaims)
	if ok && token.Valid {

	} else {
		panic(err)
	}
	jsonString, _ := json.Marshal(claims)
	fmt.Println(string(jsonString))
	var payload src.PayloadGenerate
	_ = json.Unmarshal(jsonString, &payload)

	return payload, IsExpire(payload.CreateAt)
}
