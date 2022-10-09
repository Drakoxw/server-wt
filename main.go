package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	muxer := mux.NewRouter()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}
	fmt.Printf("uri: %s\n", os.Getenv("DB_URI"))
	port = ":" + port
	log.Fatal(http.ListenAndServe(port, muxer))

}
