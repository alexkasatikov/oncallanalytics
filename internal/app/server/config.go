package server

type Config struct {
	ListenPort    string
	ListenAddress string
	DatabaseURL   string
	LogLevel      string
}

func NewConfig() *Config {
	return &Config{
		ListenPort:    "8080",
		ListenAddress: "localhost",
		DatabaseURL:   "postgres://localhost/oncallstats",
		LogLevel:      "debug",
	}
}
