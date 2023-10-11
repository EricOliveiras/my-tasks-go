package config

import "os"

type HTTPConfig struct {
	Host string
	Port string
}

func LoadHTTPConfig() HTTPConfig {
	return HTTPConfig{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
