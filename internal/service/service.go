package service

import (
	"gitnub.com/artemKapitonov/libraryAPI/internal/repository"
)

type Service struct {
	Repository
}

func New(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
