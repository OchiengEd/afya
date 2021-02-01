package api

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/OchiengEd/afya/pkg/database"
)

// Config contains microservice
// configuration
type Config struct {
	Port     string `env:"APP_PORT,default=8070"`
	Database database.Config
}

func getConfiguration() Config {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
