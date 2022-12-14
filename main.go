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
	fmt.Println(r.Header.Get("Role"))
	fmt.Fprintln(w, "Hola mundo")
}

func Req(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%q: %q\n", k, v)
	}
}

func main() {
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", Holamundo)

	muxer.HandleFunc("/api/v0/createClient/", handlers.CreateClient).Methods("POST")
	muxer.HandleFunc("/api/v0/createProvider/", handlers.CreateProvider).Methods("POST")
	muxer.HandleFunc("/api/v0/validateProvider/", handlers.ValidateAccountProvider).Methods("POST")

	muxer.HandleFunc("/api/v0/login/", handlers.Login).Methods("POST")
	muxer.HandleFunc("/api/v0/home/", middlerware.IsAuthorized(Holamundo))

	muxer.HandleFunc("/api/", Req)

	log.Fatal(http.ListenAndServe(utils.GetPort(), muxer))

}
