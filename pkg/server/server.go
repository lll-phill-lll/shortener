package server

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/lll-phill-lll/shortener/pkg/storage"
	"net/http"
	"strconv"
)

type Server interface {
	SetHandlers()
	StartServe(int) error
}

type Impl struct {
	DB     storage.DataBase
	router *mux.Router
	HostURL string
}

func (serv Impl) SetHandlers() {
	r := mux.NewRouter()
	r.HandleFunc("/short", serv.short).Methods("POST")
	r.HandleFunc("/{hash}", serv.hash).Methods("GET")
	http.Handle("/", r)
	serv.router = r

}

func (serv Impl) StartServe(port int) error {
	portStr := strconv.Itoa(port)
	portStr = ":" + portStr

	return http.ListenAndServe(portStr, serv.router)
}
