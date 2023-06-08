package main

import (
	"github.com/sirupsen/logrus"
	"gitnub.com/artemKapitonov/libraryAPI/internal/app"
)

func main() {
	app := app.New()

	if err := app.Run(); err != nil {
		logrus.Fatalf("Can't start server: %s", err.Error())
	}
 
}
