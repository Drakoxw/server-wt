package main

import (
	"fmt"
	"log"
	"net/http"
	"server/handlers"
	"server/utils"

	"github.com/gorilla/mux"
)

func main() {
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "Hola mundo") })
	muxer.HandleFunc("/api/v0/createClient/", nil)

	handlers.CreateCli()
	handlers.ValidatePass()

	log.Fatal(http.ListenAndServe(utils.GetPort(), muxer))

}
