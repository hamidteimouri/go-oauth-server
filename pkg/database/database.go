package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"log"
	"strings"
)

func Initialize(DbDriver string, DbHost string, DbPort string, DbUsername string, DbPassword string, DbName string) {
	var config string
	var err error

	validateDbDriver(DbDriver)

	/* database connection */
	if strings.ToLower(DbDriver) == "mysql" {
		config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DbUsername, DbPassword, DbHost, DbPort, DbName)

	} else if strings.ToLower(DbDriver) == "postgres" {
		config = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			DbHost, DbPort, DbUsername, DbName, DbPassword)

	}
	_, err = gorm.Open(DbDriver, config)

	if err != nil {
		fmt.Printf("Cannot connect to (%s) database", DbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("Connected to database")
	}

}

func validateDbDriver(DbDriver string) {
	driver := strings.ToLower(DbDriver)

	if DbDriver == "" {
		log.Fatalf("Database driver can not be empty")
	}

	if driver == "mysql" || driver == "postgres" {
		return
	} else {
		log.Fatalf("Database %s driver not supported", DbDriver)
	}
}
