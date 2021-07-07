package server

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallanalytics/internal/database/postgresql"
)

func opsgenieHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Opsgenie1")
	postgresql.InsertAlert()
}
