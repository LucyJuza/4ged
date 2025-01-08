package entities

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	PlayerId uint     `json:"personId"`
	Name     string   `gorm:"size:50;uniqueIndex" json:"name"`
	Password string   `json:"-"`
	ImageUrl string   `json:"image"`
	Games    []Game   `gorm:"many2many:user_games;" json:"games"`
	Plays    []Play   `gorm:"foreignkey:ID" json:"plays"`
	Players  []Player `gorm:"foreignkey:ID" json:"persons"`
}
