package middlerware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/utils"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			utils.BadResponse(w, utils.RespBad{
				Message:    "No Token Found",
				StatusCode: http.StatusForbidden,
			})
			return
		}

		var mySigningKey = []byte(utils.Secretkey)

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				utils.BadResponse(w, utils.RespBad{
					Message:    "Error in parsing token",
					StatusCode: http.StatusForbidden,
				})
				return nil, fmt.Errorf("error in parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			utils.BadResponse(w, utils.RespBad{
				Message:    "Your Token has been expired",
				StatusCode: http.StatusForbidden,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			audHttp := utils.ToStringInterface(claims["aud"])
			if !utils.ValidateOrigin(audHttp, r.Header) {
				utils.BadResponse(w, utils.RespBad{
					Message:    "Origin token no valid",
					StatusCode: http.StatusForbidden,
				})
				return
			}

			if claims["role"] != "" {
				if nivel, ok := claims["privilegeLevel"]; ok {
					r.Header.Set("privilegeLevel", utils.ToStringInterface(nivel))
				}
				r.Header.Set("Role", utils.ToStringInterface(claims["role"]))
				handler.ServeHTTP(w, r)
				return
			}
		}
		json.NewEncoder(w).Encode("Not Authorized")
	}
}
