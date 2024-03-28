package repositories

import "ozontest/database/models"

type InMemoryRepo struct {
	stf map[string]string
	fts map[string]string
}

func NewLinkMemRepo() *InMemoryRepo {
	return &InMemoryRepo{make(map[string]string), make(map[string]string)}
}

func (repo *InMemoryRepo) AddLink(links *models.Links) error {
	repo.stf[links.ShortLink] = links.FullLink
	repo.fts[links.FullLink] = links.ShortLink
	return nil
}
func (repo *InMemoryRepo) ContainsShortLink(shortLink string) bool {
	_, res := repo.stf[shortLink]
	return res
}

func (repo *InMemoryRepo) GetShortByFull(fullLink string) string {
	val, is := repo.fts[fullLink]
	if is {
		return val
	} else {
		return ""
	}
}
func (repo *InMemoryRepo) GetFullByShort(shortLink string) string {
	val, is := repo.stf[shortLink]
	if is {
		return val
	} else {
		return ""
	}
}
