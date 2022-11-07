package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"net/http"
	"os"
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

func GenerateJWT(email, user, role string, h http.Header) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["user"] = user
	claims["aud"] = Aud(h)
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateOrigin(audHttp string, h http.Header) bool {
	return audHttp == Aud(h)
}

func Aud(h http.Header) string {
	aud := ""
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	aud += host

	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			aud += ipv4.String() + ":" + addr.String() + ":"
		}
	}

	for k, v := range h {
		if k == "X-Forwarded-For" || k == "X-Envoy-External-Address" {
			aud += v[0]
		}
	}

	hash := md5.Sum([]byte(aud))
	return hex.EncodeToString(hash[:])
}
