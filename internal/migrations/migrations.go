package migrate

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func Create(db *sqlx.DB) error {

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// выполнение миграций
	if err := goose.Up(db.DB, "internal/migrations/schema"); err != nil {
		return err
	}

	return nil
}
