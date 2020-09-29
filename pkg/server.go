package pkg

import (
	"github.com/hamidteimouri/go-oauth-server/pkg/database"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct {
	DB gorm.DB
}

func Run() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error while getting ENV")
	}

	/* get env file */
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	/* database connection */
	database.Initialize(DB_DRIVER, DB_HOST, DB_PORT, DB_USERNAME, DB_PASSWORD, DB_NAME)

}
