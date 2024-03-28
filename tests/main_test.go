package main

import (
	"github.com/sirupsen/logrus"
	"ozontest/database"
	"ozontest/database/models"
	"ozontest/database/repositories"
	"testing"
)

func TestConnectDisconnect(t *testing.T) {
	testLogger := logrus.New()
	db := database.Connect(testLogger)
	_, err := db.DB()
	if err != nil {
		t.Error("Error in db.Connect()")
	}
	database.Disconnect(db, testLogger)

}

func TestNewLinkDBRepo(t *testing.T) {
	testLogger := logrus.New()
	db := database.Connect(testLogger)
	test := repositories.NewLinkDBRepo(db)
	if test.Db != db {
		t.Error("Error in creation")
	}
	database.Disconnect(db, testLogger)
}

func TestDBRepoFunctions(t *testing.T) {
	testLogger := logrus.New()
	db := database.Connect(testLogger)
	test := repositories.NewLinkDBRepo(db)
	test.AddLink(&models.Links{FullLink: "ab", ShortLink: "ahdgbcvrty"})
	if !test.ContainsShortLink("ahdgbcvrty") {
		t.Error("Error in DB function")
	}
	if test.GetShortByFull("ab") != "ahdgbcvrty" {
		t.Error("Error in DB function")
	}
	if test.GetFullByShort("ahdgbcvrty") != "ab" {
		t.Error("Error in DB function")
	}
	test.DeleteLink("ahdgbcvrty")
	database.Disconnect(db, testLogger)
}

func TestNewLinkMemRepo(t *testing.T) {
	test := repositories.NewLinkMemRepo()
	if len(test.Stf) != 0 {
		t.Error("Error in Mem creation")
	}
	if len(test.Fts) != 0 {
		t.Error("Error in Mem creation")
	}
}

func TestMemRepoFunctions(t *testing.T) {
	test := repositories.NewLinkMemRepo()
	test.AddLink(&models.Links{FullLink: "ab", ShortLink: "ahdgbcvrty"})
	if test.Stf["ahdgbcvrty"] != "ab" {
		t.Error("Error in Mem function")
	}
	if test.Fts["ab"] != "ahdgbcvrty" {
		t.Error("Error in Mem function")
	}
	if test.GetShortByFull("ab") != "ahdgbcvrty" || test.GetFullByShort("ahdgbcvrty") != "ab" {
		t.Error("Error in Mem function")
	}
}
