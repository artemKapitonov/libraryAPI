package repository

import "github.com/jmoiron/sqlx"

func NewPostgresDB() *sqlx.DB {
	return &sqlx.DB{}
}
