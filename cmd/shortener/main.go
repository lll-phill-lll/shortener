package main

import (
	"github.com/lll-phill-lll/shortener/logger"
	"os"
)

func main() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("Hello, world!")
	logger.Info.Println("Hello, world!")
}
