package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitnub.com/artemKapitonov/libraryAPI/internal/config"
	"gitnub.com/artemKapitonov/libraryAPI/internal/repository"
)

func main() {

	if err := config.Init(); err != nil {
		logrus.Fatalf("Can't init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	cfg := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := sql.Open("postgres", repository.DatabaseUrl(cfg))
	if err != nil {
		logrus.Fatalf("Can't connection with database: %s", err.Error())
	}
	logrus.Println("DB connection success")

	defer func() {
		if err := db.Close(); err != nil {
			logrus.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if err := db.Ping(); err != nil {
		logrus.Fatalf(err.Error())
	}

	if err := goose.SetDialect("postgres"); err != nil {
		logrus.Fatal("Error in set dialect")
	}

	// выполнение миграций
	if err := goose.Up(db, "internal/migrations/schema"); err != nil {
		fmt.Println("Ошибка при выполнении миграций:", err)
		return
	}
}
