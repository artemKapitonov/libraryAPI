package repository

import (
	"fmt"

	"gitnub.com/artemKapitonov/libraryAPI/internal/models"
)

func (r *Repository) NewBook(book *models.Book, userID int) (int, string, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, "", err
	}

	createBookQuery := fmt.Sprintf("INSERT INTO %s (author, title, path) VALUES ($1, $2, $3) RETURNING id", booksTable)

	var id int

	row := tx.QueryRow(createBookQuery, book.Author, book.Title, book.Path)
	if err := row.Scan(&id); err != nil {
		return 0, "", err
	}

	createUsersBookQuery := fmt.Sprintf("INSERT INTO %s (user_id, book_id) VALUES ($1, $2)", usersBooksTable)
	_, err = tx.Exec(createUsersBookQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, "", err
	}

	return id, book.Path, tx.Commit()
}

func (r *Repository) MyBooks(userID int) ([]models.BookResponse, error) {
	var books []models.BookResponse

	getMyBooksQuery := fmt.Sprintf("SELECT b.id, b.author, b.title, b.path FROM %s b INNER JOIN %s ub on b.id = ub.book_id WHERE ub.user_id = $1",
		booksTable, usersBooksTable)

	if err := r.db.Select(&books, getMyBooksQuery, userID); err != nil {
		return nil, err
	}

	return books, nil
}
