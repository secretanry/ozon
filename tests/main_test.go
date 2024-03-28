package main

import (
	"github.com/joho/godotenv"
	"ozontest/database"
	"testing"
)

func TestConnect(t *testing.T) {
	godotenv.Load()
	db := database.Connect()
	_, err := db.DB()
	if err != nil {
		t.Error("Error in db.Connect()")
	}

}
