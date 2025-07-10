package repositories

import "gorm.io/gorm"

type OrganizationRepository struct {
	DB *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{DB: db}
}
