package app

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitnub.com/artemKapitonov/libraryAPI/internal/config"
	"gitnub.com/artemKapitonov/libraryAPI/internal/handlers"
	migrate "gitnub.com/artemKapitonov/libraryAPI/internal/migrations"
	"gitnub.com/artemKapitonov/libraryAPI/internal/repository"
	"gitnub.com/artemKapitonov/libraryAPI/internal/server"
	"gitnub.com/artemKapitonov/libraryAPI/internal/service"
)

type App struct {
	repo    *repository.Repository
	service *service.Service
	handler *handlers.Handler
	server  *http.Server
}

func New() *App {

	if err := config.Init(); err != nil {
		logrus.Fatalf("Can't init configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("can't load env: %s", err.Error())
	}

	port := viper.GetString("port")

	app := &App{}

	db := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err := migrate.Create(db); err != nil {
		logrus.Fatalf("Can't do migration: %s", err.Error())
	}

	app.repo = repository.New(db)

	app.service = service.New(app.repo)

	app.handler = handlers.New(app.service)

	app.server = server.New(port, app.handler.InitRoutes())

	return app
}

func (a *App) Run() error {

	logrus.Printf("Server Listen and Serve on %s Addr", a.server.Addr)

	if err := a.server.ListenAndServe(); err != nil {
		logrus.Fatalf("Can't start server: %s", err.Error())
		return err
	}

	return nil
}
