package migrations

import (
	"gorm.io/gorm"
	"ozontest/database/models"
)

func EnableLinkAutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Links{})
	return err
}
