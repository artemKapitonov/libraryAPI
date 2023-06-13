package service

import (
	"errors"
	"fmt"
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
	AllBooks() ([]models.BookResponse, error)
	DeleteBook(bookID, userID int) error
	BookByID(bookID int) (models.BookResponse, error)
	UpdateBook(book *models.Book, userID, bookID int) error
}

func (s Service) CreateBook(book *models.Book, userID int) (int, string, error) {

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	var token = os.Getenv("YANDEX_TOKEN")

	path, err := yandex.UploadFileToYandexDisk(book, token, userID)
	if err != nil {
		return 0, "", errors.New("this book already exists")
	}

	book.Path = path

	return s.Repository.NewBook(book, userID)
}

func (s Service) MyBooks(userID int) ([]models.BookResponse, error) {
	return s.Repository.MyBooks(userID)
}

func (s Service) AllBooks() ([]models.BookResponse, error) {
	return s.Repository.AllBooks()
}

func (s Service) DeleteBook(userID, bookID int) error {

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	var token = os.Getenv("YANDEX_TOKEN")

	book, err := s.Repository.BookByID(bookID)
	if err != nil {
		return err
	}

	if err := yandex.DeleteFile(book, token, userID); err != nil {
		return err
	}

	return s.Repository.DeleteBook(bookID, userID)
}

func (s Service) BookByID(bookID int) (models.BookResponse, error) {
	return s.Repository.BookByID(bookID)
}

func (s Service) UpdateBook(book *models.Book, userID, bookID int) error {
	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	var token = os.Getenv("YANDEX_TOKEN")

	oldBook, err := s.Repository.BookByID(bookID)
	if err != nil {
		return err
	}

	if err := yandex.DeleteFile(oldBook, token, userID); err != nil {
		return err
	}

	copyPath := book.Path

	path, err := yandex.UploadFileToYandexDisk(book, token, userID)
	if err != nil {
		if err.Error() == `Put "": unsupported protocol scheme ""` {
			book.Path = copyPath
			return s.Repository.UpdateBook(book, userID, bookID)
		}
		return errors.New(fmt.Sprintf("Can't upload file to disk: %s", err.Error()))
	}

	book.Path = path

	return s.Repository.UpdateBook(book, userID, bookID)
}
