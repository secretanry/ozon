package repositories

import "ozontest/database/models"

type InMemoryRepo struct {
	Stf map[string]string
	Fts map[string]string
}

func NewLinkMemRepo() *InMemoryRepo {
	return &InMemoryRepo{make(map[string]string), make(map[string]string)}
}

func (repo *InMemoryRepo) AddLink(links *models.Links) error {
	repo.Stf[links.ShortLink] = links.FullLink
	repo.Fts[links.FullLink] = links.ShortLink
	return nil
}
func (repo *InMemoryRepo) ContainsShortLink(shortLink string) bool {
	_, res := repo.Stf[shortLink]
	return res
}

func (repo *InMemoryRepo) GetShortByFull(fullLink string) string {
	val, is := repo.Fts[fullLink]
	if is {
		return val
	} else {
		return ""
	}
}
func (repo *InMemoryRepo) GetFullByShort(shortLink string) string {
	val, is := repo.Stf[shortLink]
	if is {
		return val
	} else {
		return ""
	}
}
