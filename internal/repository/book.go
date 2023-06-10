package repository

import (
	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

func (r *Repository) NewBook(book *models.Book, userID int) (int, string, error) {
	return 0, book.Path, nil

	//query := fmt.Sprintf("INSERT INTO books (id, author, title, path) VALUES ();")
}
