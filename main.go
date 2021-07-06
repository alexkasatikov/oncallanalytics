package main

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallanalytics/pkg/webhook/alertmanager"
	"github.com/alexkasatikov/oncallanalytics/pkg/webhook/opsgenie"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/alertmanager", alertmanager.Handler)
	mux.HandleFunc("/opsgenie", opsgenie.Handler)
	log.Println("Started api server")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
