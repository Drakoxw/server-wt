package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net"
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

func GenerateJWT(email, user, role string) (string, error) {
	var mySigningKey = []byte(Secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["user"] = user
	claims["aud"] = Aud()
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Aud() string {
	aud := ""
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			aud += ipv4.String() + ":" + addr.String() + ":"
		}
	}
	aud += host
	hash := md5.Sum([]byte(aud))
	return hex.EncodeToString(hash[:])
}
