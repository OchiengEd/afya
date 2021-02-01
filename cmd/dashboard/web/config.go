package web

import (
	"log"

	env "github.com/Netflix/go-env"
)

// Config passes configuration
// to the microservice
type Config struct {
	// Port that the dashboard will listen on
	Port string `env:"PORT,default=8090"`
	// AuthProvider is the server that will authenticate users
	AuthProvider string `env:"OIDC_PROVIDER"`
	// ClientID is the ID for the Dashboard
	ClientID string `env:"CLIENT_ID"`
	// ClientSecret is the secret used to identify the dashboard
	ClientSecret string `env:"CLIENT_SECRET"`
}

func loadConfiguration() Config {
	var config Config

	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
