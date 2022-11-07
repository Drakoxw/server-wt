package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/handlers/middlerware"
	"server/services"
	"server/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func Mail(w http.ResponseWriter, r *http.Request) {
	services.SendEmail(w)
}

func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("%s uses %s\n", os.Getenv("NAME"), os.Getenv("EDITOR"))
	fmt.Printf("URI: %s\n", os.Getenv("URI"))

}

func main() {
	initEnv()
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", Holamundo)
	muxer.HandleFunc("/api/v0/createClient/", handlers.CreateClient).Methods("POST")
	muxer.HandleFunc("/api/v0/createProvider/", handlers.CreateProvider).Methods("POST")
	muxer.HandleFunc("/api/v0/login/", handlers.Login).Methods("POST")
	muxer.HandleFunc("/api/v0/home/", middlerware.IsAuthorized(Holamundo))
	muxer.HandleFunc("/api/email/", Mail).Methods("GET")
	muxer.HandleFunc("/api/", Req)

	log.Fatal(http.ListenAndServe(utils.GetPort(), muxer))

}
