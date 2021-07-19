package server

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallstats/internal/app"
	"github.com/namsral/flag"
)

func InitServer() {
	config := NewConfig()
	flag.StringVar(&config.ListenPort, "listen-port", config.ListenPort, "Port number to bind on")
	flag.StringVar(&config.ListenAddress, "listen-address", config.ListenAddress, "Address to bind on")
	flag.StringVar(&config.DatabaseURL, "database", config.DatabaseURL, "Database URL in postgres format")
	flag.StringVar(&config.LogLevel, "log-level", config.LogLevel, "Log level")
	flag.Parse()

	// Inject variables into another package
	var DatabaseURL = config.DatabaseURL
	app.DatabaseURL = DatabaseURL

	mux := http.NewServeMux()
	mux.HandleFunc("/alertmanager", app.AlertmanagerHandler)
	mux.HandleFunc("/opsgenie", app.OpsgenieHandler)

	log.Println("Server started")

	err := http.ListenAndServe(config.ListenAddress+":"+config.ListenPort, mux)
	log.Fatal(err)
}
