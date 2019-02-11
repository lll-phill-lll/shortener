package main

import (
	"github.com/lll-phill-lll/shortener/logger"
	"github.com/lll-phill-lll/shortener/pkg/application"
	"github.com/lll-phill-lll/shortener/pkg/server"
	"github.com/lll-phill-lll/shortener/pkg/storage"
	"os"
)

func GetApp() application.App {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	db := &storage.PostgresDB{Name: "links"}
	return application.App{DB: db, Server: server.Impl{DB: db, HostURL: "http://localhost:8080"}}
}

func main() {
	app := GetApp()
	err := app.DB.Init()
	if err != nil {
		logger.Error.Println(err.Error())
		os.Exit(1)
	}
	app.Server.SetHandlers()

	logger.Info.Println("Start Listening on port 8080")
	err = app.Server.StartServe(8080)
	if err != nil {
		logger.Error.Println("Can't start serving", err.Error())
	}
}
