package tokenValidator

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsValidToken(myToken string,req *http.Request) (string,bool){
	data, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		key := "THToken"
		decoded, _ :=base64.URLEncoding.DecodeString(key)
		return decoded, nil
	})
	if err != nil {
		return "",false
	}
	userId := data.Claims["Id"]

	return userId.(string),true
}
