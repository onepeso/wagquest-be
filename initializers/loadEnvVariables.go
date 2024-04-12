package initializers

import (
	"os"
)

func LoadEnvVariables() {
	// Fallback values if .env file is not present or loaded
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "default_db_host"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "default_db_port"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "default_db_user"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "default_db_password"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "default_db_name"
	}

	// Use dbHost, dbPort, dbUser, dbPassword, dbName, and other environment variables as needed
}
