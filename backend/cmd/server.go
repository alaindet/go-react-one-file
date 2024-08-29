package main

import "flag"

type ServerConfig struct {
	Port         string
	IsProduction bool
}

func ReadServerConfig() ServerConfig {
	cfg := ServerConfig{}

	flag.StringVar(&cfg.Port, "port", "8080", "The server port")
	flag.Parse()

	return cfg
}
