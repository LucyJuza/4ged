package entities

type Game struct {
	ID uint `gorm:"primaryKey" json:"id"` // not needed in post

	Name     string  `gorm:"size:300;index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE" json:"name"`
	ImageUrl string  `json:"image"`
	Genres   []Genre `gorm:"many2many:game_genres;" json:"genres"`
}
