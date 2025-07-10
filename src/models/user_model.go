package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func (User) TableName() string {
	return "users"
}
