package entities

type Player struct {
	ID     uint `gorm:"primaryKey" json:"id"` // not needed in post
	UserID uint `json:"userId"`               // not needed in post

	Name     string  `json:"name"`
	ImageUrl string  `json:"image"`
	WinRate  float32 `json:"winrate"`
}
