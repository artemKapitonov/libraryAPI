package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"gitnub.com/artemKapitonov/libraryAPI/internal/handlers"
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
	app := &App{}

	port := "8000" //TODO

	db := repository.NewPostgresDB()

	app.repo = repository.New(db)

	app.service = service.New(app.repo)

	app.handler = handlers.New(app.service)

	app.server = server.New(port, app.handler.InitRoutes())

	return app
}

func (a *App) Run() error {

	fmt.Println("server running")

	if err := a.server.ListenAndServe(); err != nil {
		logrus.Fatalf("Can't start server: %s", err.Error())
		return err
	}

	return nil
}
