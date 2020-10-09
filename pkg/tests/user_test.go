package tests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetAllUsers(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal("cannot refresh users table")
	}
	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.GetAllUsers(server.DB)

	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, len(*users), 2)
}
