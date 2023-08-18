package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
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

	getMyBooksQuery := fmt.Sprintf(`SELECT b.id, u.username, b.author, b.title, b.path FROM %s b INNER JOIN %s ub on b.id = ub.book_id
	INNER JOIN %s u on u.id = ub.user_id WHERE ub.user_id = $1`,
		booksTable, usersBooksTable, usersTable)

	if err := r.db.Select(&books, getMyBooksQuery, userID); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repository) AllBooks() ([]models.BookResponse, error) {
	var books []models.BookResponse

	getAllBooksQuery := fmt.Sprintf(`SELECT b.id, u.username, b.author, b.title, b.path FROM %s b INNER JOIN %s ub on b.id = ub.book_id
	INNER JOIN %s u on u.id = ub.user_id`, booksTable, usersBooksTable, usersTable)

	if err := r.db.Select(&books, getAllBooksQuery); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repository) BookByID(bookID int) (models.BookResponse, error) {
	var book models.BookResponse

	query := fmt.Sprintf(`SELECT b.id, u.username, b.author, b.title, b.path FROM %s b INNER JOIN %s ub on b.id = ub.book_id
	INNER JOIN %s u on u.id = ub.user_id WHERE b.id = $1`, booksTable, usersBooksTable, usersTable)

	if err := r.db.Get(&book, query, bookID); err != nil {
		return book, err
	}

	return book, nil
}

func (r *Repository) DeleteBook(bookID, userID int) error {
	deliteBookQuery := fmt.Sprintf(
		"DELETE FROM %s ub USING %s b WHERE b.id = ub.book_id AND ub.user_id = $1 AND ub.book_id = $2",
		usersBooksTable, booksTable)

	result, err := r.db.Exec(deliteBookQuery, userID, bookID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("It is not your book")
	}

	return nil
}

func (r *Repository) UpdateBook(book *models.Book, userID, bookID int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)

	argsID := 1

	if book.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argsID))
		args = append(args, book.Title)
		argsID++
	}

	if book.Author != "" {
		setValues = append(setValues, fmt.Sprintf("author=$%d", argsID))
		args = append(args, book.Author)
		argsID++
	}

	if book.Path != "" {
		setValues = append(setValues, fmt.Sprintf("path=$%d", argsID))
		args = append(args, book.Path)
		argsID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(
		"UPDATE %s b SET %s FROM %s ub WHERE b.id = ub.book_id AND ub.book_id=$%d AND ub.user_id=$%d",
		booksTable, setQuery, usersBooksTable, argsID, argsID+1)

	args = append(args, bookID, userID)
	logrus.Debugf("updateQuery: %s ", query)
	logrus.Debugf("ards: %s ", args)

	result, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("It is not your book")
	}

	return err
}
