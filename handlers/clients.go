package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/models/interfaces"
	"server/utils"
	"time"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	user := interfaces.IClienteInsert{}
	if err := bodyJson.Decode(&user); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		current_time := time.Now()
		collection, ctx, _ := db.MongoConection(w, utils.CLIENTS)
		user.Pass = utils.HashPassword(user.Pass, utils.GenerateSalt(15))
		user.Role = "user"
		user.DateReg = current_time.Format("2006-01-01 15:04:05")
		req, _ := collection.InsertOne(ctx, user)

		insertedId := req.InsertedID

		res := map[string]interface{}{
			"status":  "ok",
			"message": "Nuevo Registro",
			"data":    insertedId,
		}

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")

		if err := enc.Encode(res); err != nil {
			fmt.Println(err.Error())
		}

	}

}

func CreateCli() {
	pass := "ultrasecret"
	passEncrypt := utils.HashPassword(pass, utils.GenerateSalt(15))
	fmt.Println("passEncryp:", passEncrypt)
}
