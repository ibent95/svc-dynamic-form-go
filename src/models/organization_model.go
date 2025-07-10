package models

type Organization struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func (Organization) TableName() string {
	return "organizations"
}
