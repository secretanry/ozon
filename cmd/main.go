package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"ozontest/api/handlers"
	"ozontest/api/middleware"
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
	linksServer := handlers.NewLinksServer(logger, linkRepo)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Infoln("Failed to listen")
		logger.Debugln(err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor(logger)))
	handlers.RegisterLinksServiceServer(grpcServer, linksServer)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Infoln("failed to serve")
		logger.Debugln(err)
		return
	}
}
