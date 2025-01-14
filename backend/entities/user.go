package entities

type User struct {
	ID       uint     `gorm:"primaryKey" json:"id"`
	PlayerId uint     `json:"personId"`
	Name     string   `gorm:"size:50;uniqueIndex" json:"name"`
	Password string   `json:"-"`
	ImageUrl string   `json:"image"`
	Games    []Game   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;many2many:user_games;" json:"games"`
	Plays    []Play   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:UserID" json:"plays"`
	Players  []Player `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignkey:UserID" json:"persons"`
}

type UserCredentials struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Image    string `json:"image"`
}
