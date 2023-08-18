package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func DatabaseUrl(cfg Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Port, cfg.DBName, cfg.SSLMode)
}

func NewPostgresDB(cfg Config) *sqlx.DB {
	DBPath := DatabaseUrl(cfg)

	db, err := sqlx.Open("postgres", DBPath)
	if err != nil {
		logrus.Fatalf("Can't connection with database: %s", err.Error())
		return &sqlx.DB{}
	}

	logrus.Println("DB connection success")

	if err := db.Ping(); err != nil {
		logrus.Fatalf("Erorr Ping DB")
	}

	return db
}
