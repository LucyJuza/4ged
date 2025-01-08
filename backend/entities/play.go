package entities

type Play struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	GameId   uint     `json:"gameId"`
	Date     string   `gorm:"type:date" json:"date"`
	Location string   `json:"location"`
	Duration uint     `json:"duration"`
	Players  []Player `gorm:"foreignkey:ID" json:"participants"`
	Winners  []Player `gorm:"foreignkey:ID" json:"winners"`
}
