package models

type Publication struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func (Publication) TableName() string {
	return "publications"
}
