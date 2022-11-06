package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"server/handlers"
	"server/handlers/middlerware"
	"server/utils"

	"github.com/gorilla/mux"
)

func Holamundo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("Role"))
	fmt.Fprintln(w, "Hola mundo")
}

func ResolveHostIp() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String() + ":" + addr.String()
		}
	}
	return ""
}
func Req(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "%q: %q\n", k, v)
	}
	ip := ResolveHostIp()

	fmt.Fprintf(w, "ip: %q\n", ip)
}

func main() {
	muxer := mux.NewRouter()

	muxer.HandleFunc("/", Holamundo)
	muxer.HandleFunc("/api/v0/createClient/", handlers.CreateClient)
	muxer.HandleFunc("/api/v0/login/", handlers.Login)
	muxer.HandleFunc("/api/v0/home/", middlerware.IsAuthorized(Holamundo))
	muxer.HandleFunc("/api/", Req)

	log.Fatal(http.ListenAndServe(utils.GetPort(), muxer))

}
