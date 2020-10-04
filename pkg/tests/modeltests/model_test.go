package modeltests_test

import (
	"fmt"
	"github.com/hamidteimouri/go-oauth-server/pkg/servers"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var server = servers.Server{}

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

		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"),
			os.Getenv("TestDbHost"), os.Getenv("TestDbPort"),
			os.Getenv("TestDbName"))

		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			fmt.Printf("cannot connect to database %s", TestDbDriver)
			log.Fatal("this is error :", err)
		}

	} else if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"),
			os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_NAME"))

		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			fmt.Printf("cannot connect to database %s", TestDbDriver)
			log.Fatal("this is error :", err)
		}
	}
}

func RefreshUserTable() {
	fmt.Println("preparing...")
}
