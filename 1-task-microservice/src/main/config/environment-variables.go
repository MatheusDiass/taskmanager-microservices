package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Port             int
	DatabaseServer   string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
)

func SetupEnvironmentVariables() {
	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
	}

	//Database
	DatabaseServer = os.Getenv("DB_SERVER")
	DatabasePort, err = strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		log.Fatal(err)
	}

	DatabaseName = os.Getenv("DB_NAME")
	DatabaseUser = os.Getenv("DB_USER")
	DatabasePassword = os.Getenv("DB_PASSWORD")
}
