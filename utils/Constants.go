package utils

import "os"

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
