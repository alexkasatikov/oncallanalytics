package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/alertmanager", alertmanager.Handler)
	mux.HandleFunc("/opsgenie", server.opsgenieHandler)
	log.Println("Started api server")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
