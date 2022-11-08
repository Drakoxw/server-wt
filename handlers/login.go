package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/models/interfaces"
	"server/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

func LoginOld(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	dataLogin := interfaces.Login{}
	if err := bodyJson.Decode(&dataLogin); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		passEncryp := utils.HashPassword(dataLogin.Pass, utils.GenerateSalt(15))
		dataLogin.Pass = passEncryp
		collection, ctx, _ := db.MongoConection(w, utils.CLIENTS)
		result := interfaces.IClienteInsert{}
		err := collection.FindOne(ctx, dataLogin).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				utils.BadResponse(w, utils.RespBad{
					Message:    "Error en el correo o contraseña",
					StatusCode: http.StatusNotFound,
				})
			}
		}
		token, _ := utils.GenerateJWT(result.Email, result.UserName, result.Role, r.Header)
		response := fmt.Sprintf(`{
				"token": "%s"
			}`, token)
		utils.ResponseFile(w, string(response), "usuario validado")
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	dataLogin := interfaces.Login{}
	if err := bodyJson.Decode(&dataLogin); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		passEncryp := utils.HashPassword(dataLogin.Pass, utils.GenerateSalt(15))
		dataLogin.Pass = passEncryp
		token := ""
		canalCli := make(chan string)
		canalPro := make(chan string)

		go func() { canalCli <- searchClient(w, r, dataLogin) }()
		go func() { canalPro <- searchProvider(w, r, dataLogin) }()

		if token == "" {
			token = <-canalCli
		}
		if token == "" {
			token = <-canalPro
		}

		if token == "" {
			utils.BadResponse(w, utils.RespBad{
				Message:    "Error en el correo o contraseña",
				StatusCode: http.StatusNotFound,
			})
			return
		}
		response := fmt.Sprintf(`{
				"token": "%s"
			}`, token)
		utils.ResponseFile(w, string(response), "usuario validado")
	}
}

func searchProvider(w http.ResponseWriter, r *http.Request, dataLogin interfaces.Login) string {
	collection, ctx, _ := db.MongoConection(w, utils.PROVIDERS)
	result := interfaces.IProviderInsert{}
	err := collection.FindOne(ctx, dataLogin).Decode(&result)
	if err != nil {
		return ""
	}
	token, _ := utils.GenerateJWT(result.Email, result.UserName, result.Role, r.Header)
	return token
}

func searchClient(w http.ResponseWriter, r *http.Request, dataLogin interfaces.Login) string {
	collection, ctx, _ := db.MongoConection(w, utils.CLIENTS)
	result := interfaces.IClienteInsert{}
	err := collection.FindOne(ctx, dataLogin).Decode(&result)
	if err != nil {
		return ""
	}
	token, _ := utils.GenerateJWT(result.Email, result.UserName, result.Role, r.Header)
	return token
}
