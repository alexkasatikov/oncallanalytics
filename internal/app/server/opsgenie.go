package server

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallstats/internal/database/postgresql"
)

func opsgenieHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Opsgenie1")
	postgresql.InsertAlert()
}
