package bootstrap

import (
	"github.com/hamidteimouri/go-oauth-server/pkg/database"
	"github.com/hamidteimouri/go-oauth-server/pkg/servers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = servers.Server{}

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

	server.Initialize()

	/* run server */
	server.RunningReport("127.0.0.1:8080")

}
