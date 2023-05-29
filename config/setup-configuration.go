package config

import (
	"os"

	"github.com/joho/godotenv"
)

type SetupConfiguration struct {
	dbName, dbPassword, dbUsername, dbHost, dbPort string
}

func Configuration() SetupConfiguration {
	var err = godotenv.Load()

	if err != nil {
		panic(err.Error())
	}

	return SetupConfiguration{
		dbName:     os.Getenv("DB_NAME"),
		dbPassword: os.Getenv("DB_PASS"),
		dbHost:     os.Getenv("DB_HOST"),
		dbUsername: os.Getenv("DB_USERNAME"),
		dbPort:     os.Getenv("DB_PORT"),
	}
}
