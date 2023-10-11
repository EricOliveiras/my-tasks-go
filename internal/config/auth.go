package config

import "os"

type AuthConfig struct {
	AccessSecret string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret: os.Getenv("SECRET"),
	}
}
