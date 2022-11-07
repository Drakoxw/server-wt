package db

import (
	"context"
	"fmt"
	"net/http"
	"server/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConection(w http.ResponseWriter, nameColl string) (*mongo.Collection, context.Context, error) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.UriVar))
	if err != nil {
		fmt.Println(err.Error())
	}
	DB := client.Database(utils.MONGO_DB)
	collection := DB.Collection(nameColl)

	return collection, ctx, err
}
