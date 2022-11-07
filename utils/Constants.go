package utils

import "os"

var UriVar = "mongodb+srv://DrakoMaster:xnkKnbYuGbGVegXP@drako-db.fguhd.mongodb.net/?retryWrites=true&w=majority"

const (
	MONGO_DB  = "Go_Test"
	CLIENTS   = "Clients"
	PROVIDERS = "Providers"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}
	return ":" + port
}
