package application

import (
	"github.com/lll-phill-lll/shortener/pkg/server"
	"github.com/lll-phill-lll/shortener/pkg/storage"
)

type App struct {
	DB     storage.DataBase
	Server server.Impl
	HostURL string
}
