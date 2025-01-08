package entities

import (
	"log"

	"gorm.io/gorm"
)

type Game struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"size:300;index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE" json:"name"`
	ImageUrl string  `json:"image"`
	Genres   []Genre `gorm:"many2many:game_genres;" json:"genres"`
}

func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println(g.Genres)

	return
}
