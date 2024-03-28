package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func Connect() *gorm.DB {
	customLogger := logger.New(
		log.New(nil, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Silent,
		},
	)

	err := godotenv.Load()
	if err != nil {
		logrus.Infoln("Cannot load environment variables")
		logrus.Debugln(err)
		return nil
	}
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, dbname, port)
	_, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: customLogger})
	if err != nil {
		logrus.Infoln("Cannot open database, attempting to create")
		dsnTemp := fmt.Sprintf("host=%s user=postgres password=%s dbname=postgres port=%s sslmode=disable",
			host, os.Getenv("POSTGRES_USER_PASSWORD"), port)
		tempDb, err := gorm.Open(postgres.Open(dsnTemp), &gorm.Config{Logger: customLogger})
		if err != nil {
			logrus.Infof("Cannot create new database, shutting down %s", err)
			logrus.Debugln(err)
			return nil
		}
		query := fmt.Sprintf("CREATE DATABASE %s", dbname)
		err = tempDb.Exec(query).Error
		if err != nil {
			logrus.Infoln("Cannot create new database, shutting down")
			logrus.Debugln(err)
			return nil
		}
		query = fmt.Sprintf("CREATE USER %s WITH PASSWORD '%s'", username, password)
		err = tempDb.Exec(query).Error
		if err != nil {
			logrus.Infoln("Cannot create a user, shutting down")
			logrus.Debugln(err)
			return nil
		}
		query = fmt.Sprintf("ALTER DATABASE %s OWNER TO %s", dbname, username)
		err = tempDb.Exec(query).Error
		if err != nil {
			logrus.Infoln("Cannot alter database, shutting down")
			logrus.Debugln(err)
			return nil
		}
		dbSQL, err := tempDb.DB()
		if err != nil {
			logrus.Infoln("Cannot close temporary connection")
			logrus.Debugln(err)
			return nil
		}
		err = dbSQL.Close()
		if err != nil {
			logrus.Infoln("Cannot close temporary connection")
			logrus.Debugln(err)
			return nil
		}
		logrus.Infoln("Database successfully created!")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: customLogger})
	if err != nil {
		logrus.Infoln("Cannot reconnect to database, shutting down")
		logrus.Debugln(err)
		return nil
	}
	return db
}

func Disconnect(db *gorm.DB) {
	temp, err := db.DB()
	if err != nil {
		logrus.Infoln("Disconnect error")
		logrus.Debugln(err)
		return
	}
	err = temp.Close()
	if err != nil {
		logrus.Infoln("Disconnect error")
		logrus.Debugln(err)
		return
	}

}
