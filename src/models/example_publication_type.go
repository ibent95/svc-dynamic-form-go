package models

import "time"

type ExamplePublicationType struct {
	ID uint `gorm:"primaryKey;column:id;not null" json:"id"`
	PublicationTypeCode string `gorm:"size:5;column:publication_type_code;not null" json:"publication_type_code"`
	PublicationTypeName string `gorm:"size:255;column:publication_type_name;not null" json:"publication_type_name"`
	Description string `gorm:"type:text;column:description" json:"description"`
	ActiveFlag bool `gorm:"column:active_flag;not null;default:true" json:"active_flag"`
	UserCreate string `gorm:"size:150;column:user_create;not null;default:'system'" json:"user_create"`
	CreatedDatetime time.Time `gorm:"column:created_datetime:not null;autoCreateTime" json:"created_datetime"`
	UserUpdate string `gorm:"size:150;column:user_update;not null;default:'system'" json:"user_update"`
	UpdatedDatetime time.Time `gorm:"column:updated_datetime:not null;autoUpdateTime" json:"updated_datetime"`

	Publications []ExamplePublication `gorm:"foreignKey:PublicationTypeId" json:"publications"`
}

func (ExamplePublicationType) TableName() string {
	return "example_publication_type"
}
