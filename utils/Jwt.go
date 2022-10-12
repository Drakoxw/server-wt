package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	Role        string `json:"role"`
	User        string `json:"user"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

var Secretkey = "UltraSecret"

func GenerateJWT(email, user, role string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		// fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
