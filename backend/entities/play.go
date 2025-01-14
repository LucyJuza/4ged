package entities

import "time"

type Play struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	UserID   uint      `json:"userId"`
	GameId   uint      `json:"gameId"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
	Duration uint      `json:"duration"`
	Players  []Player  `gorm:"many2many:play_participants;"  json:"participants"`
	Winners  []Player  `gorm:"many2many:play_winners;"  json:"winners"`
}
