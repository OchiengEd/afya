package database

import (
	"fmt"
	"os"
)

// Schema interface defines characters
// of a database schema
type Schema interface {
	Definition() string
	Name() string
}

// Initialize defines database schema
func Initialize(db Schema) {}

// Connect initiates connection to database
func Connect() {
	host := os.Getenv("DATABASE_HOST")
	port := requireEnv("DATABASE_PORT", "5432")
	database := requireEnv("DATABASE_NAME", "afya")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")

	_ = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
}

func requireEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}

	return fallback
}

func ifDBExists(db Schema) {}

func ifTableExists(db Schema) {}
