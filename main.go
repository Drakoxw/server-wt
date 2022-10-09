package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const URI = "mongodb+srv://DrakoMaster:xnkKnbYuGbGVegXP@drako-db.fguhd.mongodb.net/?retryWrites=true&w=majority"

func main() {
	muxer := mux.NewRouter()

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
