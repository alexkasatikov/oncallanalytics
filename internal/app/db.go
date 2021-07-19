package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateAlerts(dsn string, alert Alert) uint64 {

	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var id uint64

	if alert.Status == "firing" {
		query := "INSERT INTO alerts (fingerprint, status, startsat, endsat) VALUES ($1, $2, $3, $4) RETURNING id"
		row := conn.QueryRow(context.Background(),
			query,
			alert.Fingerprint, alert.Status, alert.StartsAt, alert.EndsAt)

		err = row.Scan(&id)

		if err != nil {
			query := "SELECT id FROM alerts WHERE fingerprint = $1"
			row := conn.QueryRow(context.Background(),
				query,
				alert.Fingerprint)

			err = row.Scan(&id)

			if err != nil {
				log.Printf("Unable to find alert in database: %v\n", err)
			}
		}
	} else {
		query := "UPDATE alerts SET (status, endsat) = ($1, $2) WHERE fingerprint = $3 RETURNING id"
		row := conn.QueryRow(context.Background(),
			query,
			alert.Status, alert.EndsAt, alert.Fingerprint)

		err = row.Scan(&id)

		if err != nil {
			log.Printf("Unable to UPDATE alert: %v\n", err)
			id = 0
		}
	}

	return id
}

func UpdateLabels(dsn string, labels map[string]string) []uint64 {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	query := "INSERT INTO labels (key, value) VALUES ($1, $2) RETURNING id"
	var labelsIds []uint64
	for key, val := range labels {
		labelsRow := conn.QueryRow(context.Background(),
			query,
			key, val)

		var id uint64
		err = labelsRow.Scan(&id)

		if err != nil {
			query := "SELECT id FROM labels WHERE key = $1 and value = $2"
			row := conn.QueryRow(context.Background(),
				query,
				key, val)

			err = row.Scan(&id)

			if err != nil {
				log.Printf("Unable to find label in database: %v\n", err)
			}
		}
		labelsIds = append(labelsIds, id)
	}
	return labelsIds
}

func UpdateAlertsLabels(dsn string, alertId uint64, labelsIds []uint64) {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	query := "INSERT INTO alerts_labels (alert_id, label_id) VALUES ($1, $2) RETURNING id"
	for _, val := range labelsIds {
		row := conn.QueryRow(context.Background(),
			query,
			alertId, val)

		var id uint64
		err = row.Scan(&id)
		if err != nil {
			//log.Printf("Unable to INSERT alerts_labels: %v\n", err)
			//id = 0
		}
	}
}
