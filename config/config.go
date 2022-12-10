package config

import (
	"final_project_4/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverPort := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbName := os.Getenv("MYSQLDATABASE")

	config := models.Config{
		ServerPort: serverPort,
		Database: models.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
		},
	}
	return config
}
