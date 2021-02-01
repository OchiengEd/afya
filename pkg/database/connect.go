package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Netflix/go-env"
)

// Schema interface defines characters
// of a database schema
type Schema interface {
	Definition() string
	Name() string
}

// Config reads database
// parameters from env variables
type Config struct {
	Port         string `env:"DB_PORT,default=5432"`
	Host         string `env:"DB_HOST"`
	Username     string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	PasswordFile string `env:"DB_PASSWORD_FILE"`
	Database     string `env:"DB_NAME"`
}

// Configuration returns database
// configuration from environment variables
func getConfiguration() Config {
	var config Config
	_, err := env.UnmarshalFromEnviron(&config)
	if err != nil {
		log.Fatal(err)
	}

	if config.PasswordFile != "" {
		var password []byte
		password, err = ioutil.ReadFile(config.PasswordFile)
		if err != nil {
			log.Fatal(err)
		}

		config.Password = string(password)
	}

	return config
}

// Initialize defines database schema
func Initialize(db Schema) {}

// Connect initiates connection to database
func Connect() {
	config := getConfiguration()

	_ = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username, config.Password, config.Database)
}

func requireEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return fallback
}
