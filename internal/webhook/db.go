package webhook

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateAlerts(dsn string, alert Alert) {

	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	if alert.Status == "firing" {
		query := "INSERT INTO alerts (fingerprint, status, startsat, endsat) VALUES ($1, $2, $3, $4) RETURNING id"
		row := conn.QueryRow(context.Background(),
			query,
			alert.Fingerprint, alert.Status, alert.StartsAt, alert.EndsAt)

		var id uint64
		err = row.Scan(&id)

		if err != nil {
			log.Printf("Unable to INSERT new alert: %v\n", err)
			return
		}
	} else {
		query := "UPDATE alerts SET (status, endsat) = ($1, $2) WHERE fingerprint = $3 RETURNING id"
		row := conn.QueryRow(context.Background(),
			query,
			alert.Status, alert.EndsAt, alert.Fingerprint)

		var id uint64
		err = row.Scan(&id)

		if err != nil {
			log.Printf("Unable to UPDATE alert: %v\n", err)
			return
		}
	}
}
