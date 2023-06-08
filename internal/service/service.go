package service

import (
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
	"gitnub.com/artemKapitonov/libraryAPI/internal/repository"
)

type Service struct {
	Repository
}

type Repository interface {
	NewUser(user models.User) (int, error)
	User(username, password string) (models.User, error)
	NewBook(input models.Book, userID int) (int, error)
}

func New(repo *repository.Repository) *Service {
	return &Service{Repository: repo}
}
