package repositories

import "gorm.io/gorm"

type PublicationRepository struct {
	DB *gorm.DB
}

func NewPublicationRepository(db *gorm.DB) *PublicationRepository {
	return &PublicationRepository{DB: db}
}
