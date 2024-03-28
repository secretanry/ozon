package repositories

import "ozontest/database/models"

type Repo interface {
	AddLink(links *models.Links) error
	ContainsShortLink(shortLink string) bool
	GetShortByFull(fullLink string) string
	GetFullByShort(shortLink string) string
}
