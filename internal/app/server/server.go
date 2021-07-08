package server

import (
	"log"
	"net/http"

	"github.com/alexkasatikov/oncallstats/internal/webhook"
	"github.com/namsral/flag"
)

func InitServer() {
	config := NewConfig()
	flag.StringVar(&config.ListenPort, "listen-port", config.ListenPort, "Port number to bind on")
	flag.StringVar(&config.ListenAddress, "listen-address", config.ListenAddress, "Address to bind on")
	flag.StringVar(&config.DatabaseURL, "database", config.DatabaseURL, "Database URL in postgres format")
	flag.StringVar(&config.LogLevel, "log-level", config.LogLevel, "Log level")

	flag.Parse()
	mux := http.NewServeMux()
	//mux.HandleFunc("/alertmanager", AlertmanagerHandler)
	mux.HandleFunc("/opsgenie", webhook.OpsgenieHandler)

	err := http.ListenAndServe(config.ListenAddress+":"+config.ListenPort, mux)
	log.Fatal(err)
}
