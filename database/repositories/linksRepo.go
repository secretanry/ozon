package repositories

import (
	"gorm.io/gorm"
	"ozontest/database/models"
)

type LinksRepo struct {
	db *gorm.DB
}

func NewLinkDBRepo(db *gorm.DB) *LinksRepo {
	return &LinksRepo{db: db}
}

func (repo *LinksRepo) AddLink(links *models.Links) error {
	return repo.db.Create(links).Error
}

func (repo *LinksRepo) ContainsShortLink(shortLink string) bool {
	err := repo.db.Where("short_link=?", shortLink).First(&models.Links{}).Error
	if err == nil {
		return true
	}
	return false
}

func (repo *LinksRepo) GetShortByFull(fullLink string) string {
	var entity models.Links
	err := repo.db.Where("full_link=?", fullLink).First(&entity).Error
	if err != nil {
		return ""
	}
	return entity.ShortLink
}

func (repo *LinksRepo) GetFullByShort(shortLink string) string {
	var entity models.Links
	err := repo.db.Where("short_link=?", shortLink).First(&entity).Error
	if err != nil {
		return ""
	}
	return entity.FullLink
}
