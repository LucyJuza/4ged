package entities

type Genre struct {
	Name string `gorm:"primaryKey" json:"name"`
}
