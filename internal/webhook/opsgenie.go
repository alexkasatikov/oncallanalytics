package webhook

import (
	"net/http"

	"github.com/alexkasatikov/oncallstats/internal/database/postgresql"
)

var DatabaseURL string

func OpsgenieHandler(w http.ResponseWriter, r *http.Request) {
	postgresql.InsertAlert(DatabaseURL)
}
