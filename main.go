package main

import (
	"hirasawa/bancho/server"
	"log"
	"net/http"
)

func main() {
	bancho := server.New()
	banchoServer := &http.Server{
		Addr:    ":49152",
		Handler: bancho,
	}

	log.Fatal(banchoServer.ListenAndServe())
}
