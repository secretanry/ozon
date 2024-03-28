package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"ozontest/api/handlers"
	"ozontest/api/middleware"
	"ozontest/api/payload/responses"
	"ozontest/database"
	"ozontest/database/migrations"
	"ozontest/database/repositories"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot parse environment variables")
	}

	logger := logrus.New()
	if os.Getenv("LOGS_FORMAT") == "JSON" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	}
	switch os.Getenv("LOGS_LEVEL") {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
		break
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(logger.Writer())
	var linkRepo repositories.Repo
	if os.Getenv("STORAGE") == "DB" {
		db := database.Connect(logger)
		err = migrations.EnableLinkAutoMigration(db)
		if err != nil {
			logger.Infoln("Cannot enable auto-migration")
			logger.Debugln(err)
			return
		}
		linkRepo = repositories.NewLinkDBRepo(db)
		defer database.Disconnect(db, logger)
	} else {
		linkRepo = repositories.NewLinkMemRepo()
	}

	http.HandleFunc("/link", middleware.LoggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.LongHandler(w, r, linkRepo, logger)
		case http.MethodGet:
			handlers.ShortLinkHandler(w, r, linkRepo, logger)
		default:
			logrus.Infoln("Method not allowed")
			w.WriteHeader(405)
			jsonResponse, err := json.Marshal(&responses.RespModel{Message: "Method not allowed"})
			if err != nil {
				logrus.Infoln("Cannot marshall json")
				logrus.Debugln(err)
				return
			}
			_, err = w.Write(jsonResponse)
			if err != nil {
				logrus.Infoln("Cannot send a response")
				logrus.Debugln(err)
				return
			}
		}
	}, logger))
	logrus.Infoln(http.ListenAndServe(":8080", nil))
}
