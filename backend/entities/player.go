package entities

type Player struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	UserID   uint    `json:"userId"`
	Name     string  `json:"name"`
	ImageUrl string  `json:"image"`
	WinRate  float32 `json:"winrate"`
}
