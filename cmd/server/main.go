package main

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallstats/internal/app/server"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/alertmanager", server.AlertmanagerHandler)
	mux.HandleFunc("/opsgenie", server.OpsgenieHandler)
	log.Println("Started api server")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
