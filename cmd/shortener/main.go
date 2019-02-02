package main

import (
	"github.com/gorilla/mux"
	"github.com/lll-phill-lll/shortener/logger"
	"github.com/lll-phill-lll/shortener/pkg/server"
	"net/http"
	"os"
)

func InitApp() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func main() {
	InitApp()
	r := mux.NewRouter()
	r.HandleFunc("/short", server.Short).Methods("POST")
	r.HandleFunc("/{hash}", server.Hash).Methods("GET")
	http.Handle("/", r)

	logger.Info.Println("Start Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error.Println(err.Error())
	}
}
