package postgresql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertAlert() {
	//conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	conn, err := pgxpool.Connect(context.Background(), "postgres://localhost/seshat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	row := conn.QueryRow(context.Background(),
		"INSERT INTO phonebook (name) VALUES ($1) RETURNING id",
		"test")

	var id uint64
	err = row.Scan(&id)

	fmt.Println("111")
}
