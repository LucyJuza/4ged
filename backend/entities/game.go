package entities

type Game struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"uniqueIndex" json:"name"`
	ImageUrl string  `json:"image"`
	Genres   []Genre `gorm:"many2many:game_genres;" json:"genres"`
}
