package entities

type Game struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"size:300;index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE" json:"name"`
	ImageUrl string  `json:"image"`
	Genres   []Genre `gorm:"many2many:game_genres;" json:"genres"`
}

type UGame struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	UserID   uint    `json:"userId"`
	Name     string  `gorm:"size:300;index:,class:FULLTEXT,option:WITH PARSER ngram INVISIBLE" json:"name"`
	ImageUrl string  `json:"image"`
	Genres   []Genre `gorm:"many2many:ugame_genres;" json:"genres"`
}
