package service

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
	"gitnub.com/artemKapitonov/libraryAPI/internal/service/yandex"
)

type Repository interface {
	NewUser(user models.User) (int, error)
	User(username, password string) (models.User, error)
	NewBook(book *models.Book, userID int) (int, string, error)
	MyBooks(userID int) ([]models.BookResponse, error)
}

func (s Service) CreateBook(book *models.Book, userID int) (int, string, error) {

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	var token = os.Getenv("YANDEX_TOKEN")

	path, err := yandex.UploadFileToYandexDisk(book, token)
	if err != nil {
		return 0, "", err
	}

	book.Path = path

	return s.Repository.NewBook(book, userID)
}

func (s Service) MyBooks(userID int) ([]models.BookResponse, error) {
	return s.Repository.MyBooks(userID)
}
