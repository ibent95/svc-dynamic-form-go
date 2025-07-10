package services

import "svc-dynamic-form-go/src/repositories"

type OrganizationService struct {
	Repo *repositories.OrganizationRepository
}

func NewOrganizationService(repo *repositories.OrganizationRepository) *OrganizationService {
	return &OrganizationService{Repo: repo}
}
