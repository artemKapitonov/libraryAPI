package service

import "gitnub.com/artemKapitonov/libraryAPI/internal/models"

func (s Service) CreateBook(input models.Book, userID int) (int, error) {
	return s.Repository.NewBook(input, userID)
}
