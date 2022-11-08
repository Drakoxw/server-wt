package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/db"
	"server/models/interfaces"
	services "server/services/emails"
	"server/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProvider(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	user := interfaces.IProviderInsert{}
	if err := bodyJson.Decode(&user); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
	} else {
		current_time := time.Now()
		collection, ctx, _ := db.MongoConection(w, utils.PROVIDERS)
		user.Pass = utils.HashPassword(user.Pass, utils.GenerateSalt(15))
		user.Role = "provider"
		user.PrivilegeLevel = 5
		user.Verify = false
		user.DateReg = current_time.Format("2006-01-01 15:04:05")
		req, _ := collection.InsertOne(ctx, user)

		insertedId := req.InsertedID

		res := map[string]interface{}{
			"status":  "ok",
			"message": "Nuevo Registro",
			"data":    insertedId,
		}

		go services.SendWelcomeProvider(
			user.UserName,
			user.Email,
			utils.ToStringIdPrimite(req.InsertedID),
			user.NameEnterprise)

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")

		if err := enc.Encode(res); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func ValidateAccountProvider(w http.ResponseWriter, r *http.Request) {
	bodyJson := json.NewDecoder(r.Body)
	userID := interfaces.IdValidate{}
	if err := bodyJson.Decode(&userID); err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    "No se pudo procesar el recurso",
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}
	collection, ctx, _ := db.MongoConection(w, utils.PROVIDERS)
	id, _ := primitive.ObjectIDFromHex(userID.Id)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verify", Value: true}}}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.BadResponse(w, utils.RespBad{
			Message:    err.Error(),
			StatusCode: http.StatusUnprocessableEntity,
		})
		return
	}
	getProvi := interfaces.IProviderInsert{}
	err = collection.FindOne(ctx, filter).Decode(&getProvi)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.BadResponse(w, utils.RespBad{
				Message:    err.Error(),
				StatusCode: http.StatusNotFound,
			})
		}
	}
	token, _ := utils.GenerateJWT(getProvi.Email, getProvi.UserName, getProvi.Role, r.Header)
	response := fmt.Sprintf(`{
		"token": "%s"
	}`, token)
	utils.ResponseFile(w, string(response), "usuario validado")
}
