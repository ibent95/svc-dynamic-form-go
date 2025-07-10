package models

import "time"

type ExamplePublication struct {
	ID uint `gorm:"column:id;not null;primaryKey;autoIncrement:false" json:"id"` // UUID_SHORT()
	UUID string `gorm:"column:uuid;type:char(36);uniqueIndex" json:"uuid"`
	PublicationTypeId uint `gorm:"column:publication_type_id;not null" json:"publication_type_id"`
	Title string `gorm:"size:100;column:judul;not null" json:"judul"`
	PublishedFlag bool `gorm:"column:published_flag;not null;default:false" json:"published_flag"`
	PublishDate time.Time `gorm:"type:date;column:publish_date" json:"publish_date"`
	ActiveFlag bool `gorm:"column:active_flag;not null;default:true" json:"active_flag"`
	UserCreate string `gorm:"size:150;column:user_create;not null;default:'system'" json:"user_create"`
	CreatedDatetime time.Time `gorm:"column:created_datetime:not null;autoCreateTime" json:"created_datetime"`
	UserUpdate string `gorm:"size:150;column:user_update;not null;default:'system'" json:"user_update"`
	UpdatedDatetime time.Time `gorm:"column:updated_datetime:not null;autoUpdateTime" json:"updated_datetime"`

	PublicationType *ExamplePublicationType `gorm:"foreignKey:PublicationTypeId" json:"publication_type"`
}

func (ExamplePublication) TableName() string {
	return "example_publication"
}
