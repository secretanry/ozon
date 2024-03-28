package models

type Links struct {
	ID        uint
	ShortLink string `gorm:"primaryKey" gorm:"index" gorm:"serial"`
	FullLink  string
}
