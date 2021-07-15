package webhook

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertNewAlert(database string, alert Alert) {
	log.Println("Inserting new alert")
	conn, err := pgxpool.Connect(context.Background(), database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO alerts (fingerprint, status, startsat, endsat) VALUES ($1, $2, $3, $4) RETURNING id",
		alert.Fingerprint, alert.Status, alert.StartsAt, alert.EndsAt)

	var id uint64
	err = row.Scan(&id)

	if err != nil {
		log.Printf("Unable to INSERT new alert: %v\n", err)
		return
	}

	fmt.Println("done")
}

func ResolveAlert(database string, alert Alert) {
	log.Println("Inserting new alert")
	conn, err := pgxpool.Connect(context.Background(), database)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	row := conn.QueryRow(context.Background(),
		"UPDATE alerts SET (status, endsat) = ($1, $2) WHERE fingerprint = $3 RETURNING id",
		alert.Status, alert.EndsAt, alert.Fingerprint)

	var id uint64
	err = row.Scan(&id)

	if err != nil {
		log.Printf("Unable to INSERT new alert: %v\n", err)
		return
	}

	fmt.Println("done")
}
