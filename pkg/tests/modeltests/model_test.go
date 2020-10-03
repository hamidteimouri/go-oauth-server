package modeltests_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../../.env"))

	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")

	if TestDbDriver == "mysql" {

		BDURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"),
			os.Getenv("TestDbHost"), os.Getenv("TestDbPort"),
			os.Getenv("TestDbName"))

	}
	if TestDbDriver == "postgres" {

	}
}
