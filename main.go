package main

import (
	"fmt"
	"hirasawa/bancho"
	"log"
	"net/http"
)

func handleBancho(w http.ResponseWriter, r *http.Request) {
	if r.UserAgent() != "osu!" && r.Method == "GET" {
		fmt.Fprint(w, "Fuck you")
		return
	}

	if r.Method == "POST" {
		if osu_token := r.Header.Get("Osu-Token"); osu_token != "" {
			log.Println("Bancho request sent by", osu_token)
		    bancho.HandleBanchoRequest(w, r)
		} else {
			log.Println("Bancho login request from", r.UserAgent())
			bancho.HandleBanchoLogin(w, r)
		}
	}
}

func main() {
	banchoMux := http.NewServeMux()
	banchoMux.HandleFunc("/", handleBancho)

	banchoServer := &http.Server{
		Addr:    ":49152",
		Handler: banchoMux,
	}

	log.Fatal(banchoServer.ListenAndServe())
}
