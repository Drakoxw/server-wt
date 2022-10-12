package main

import (
	"fmt"
	"log"
	"net/http"
	"server/handlers"
	"server/handlers/middlerware"
	"server/utils"

	"github.com/gorilla/mux"
)

func Holamundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola mundo")
}

func main() {
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", Holamundo)
	muxer.HandleFunc("/api/v0/createClient/", handlers.CreateClient)
	muxer.HandleFunc("/api/v0/login/", handlers.Login)
	muxer.HandleFunc("/api/v0/home/", middlerware.IsAuthorized(Holamundo))

	log.Fatal(http.ListenAndServe(utils.GetPort(), muxer))

}
