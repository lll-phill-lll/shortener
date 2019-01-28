package main

import (
	"github.com/gorilla/mux"
	"github.com/lll-phill-lll/shortener/logger"
	"net/http"
	"os"
)

func InitApp() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func short(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("short")
}


func hash(w http.ResponseWriter, r *http.Request) {
	logger.Debug.Println("hash")
}


func main() {
	InitApp()
	r := mux.NewRouter()
	r.HandleFunc("/short", short).Methods("POST")
	r.HandleFunc("/{hash}", hash).Methods("GET")
	http.Handle("/", r)

	logger.Info.Println("Start Listening on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error.Println(err.Error())
	}
}
