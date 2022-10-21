package main

import (
	"log"
	"os"
	"testing"

	"github.com/JohnBurtt10/go/app/models"
	"github.com/JohnBurtt10/go/app/repos"
	"github.com/JohnBurtt10/go/app/utils/password/argon2id"
	"github.com/JohnBurtt10/go/database"
)

func setup() {
	argon2ID := argon2id.New()
	hash, err := argon2ID.GenerateFromPassword("correctpassword")

	if err != nil {
		panic(err)
	}
	//create a test user
	testUser := &models.User{
		Username: "Test",
		Password: hash,
	}

	// Create a user, if error return
	if err := repos.CreateUser(testUser); err != nil {
		log.Printf("error creating test user: %s", err.Error())
	}

}

func shutdown() {
	testUser := models.User{}
	database.DBConn.Find(&testUser, "username = ?", "test")

	if err := database.DBConn.Unscoped().Delete(&testUser); err.Error != nil {
		log.Printf("error deleting test user: %s", err.Error)
	}
	database.DBConn.Unscoped().Delete(testUser)

}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
