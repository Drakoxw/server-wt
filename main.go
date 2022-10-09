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

	URI := os.Getenv("DB_URI")

	muxer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, URI)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3300"
	}
	fmt.Printf("uri: %s\n", URI)
	port = ":" + port
	log.Fatal(http.ListenAndServe(port, muxer))

}
