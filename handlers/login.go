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
		collection, ctx, _ := db.MongoConection(w, utils.CLIENTS)
		result := interfaces.IClienteInsert{}
		dataLogin.Pass = passEncryp
		err := collection.FindOne(ctx, dataLogin).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				utils.BadResponse(w, utils.RespBad{
					Message:    "Error un correo o contraseña",
					StatusCode: http.StatusNotFound,
				})
			}
		}
		token, _ := utils.GenerateJWT(result.Email, result.UserName, result.Role, r.Header)
		response := fmt.Sprintf(`{
				"token": "%s"
			}`, token)
		utils.ResponseFile(w, string(response))
	}
}
