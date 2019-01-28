package main

import (
	"github.com/lll-phill-lll/shortener/logger"
	"os"
)

func InitApp() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

}

func main() {
	logger.Info.Println("Hello, world!")
	logger.Info.Println("Hello, world!")
}
