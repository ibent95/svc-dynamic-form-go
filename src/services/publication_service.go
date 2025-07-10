package services

import "svc-dynamic-form-go/src/repositories"

type PublicationService struct {
	Repo *repositories.PublicationRepository
}

func NewPublicationService(repo *repositories.PublicationRepository) *PublicationService {
	return &PublicationService{Repo: repo}
}
