package server

import (
	"log"

	"github.com/namsral/flag"
)

type Config struct {
	BindPort int
}

func NewConfig() Config {
	config := Config{}
	config.BindPort = 8080
	return config
}

func InitServer() {
	config := NewConfig()
	flag.IntVar(&config.BindPort, "port", config.BindPort, "Port number")
	flag.Parse()
	log.Println("You seem to prefer", config.BindPort)
}
