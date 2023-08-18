package repository

import "github.com/jmoiron/sqlx"

const (
	usersTable      = "users"
	booksTable      = "books"
	usersBooksTable = "user_books"
)

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}
