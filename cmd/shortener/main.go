package main

import (
	"os"
	"shortener/logger"
)

func main() {
	logger.SetLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	logger.Info.Println("Hello, world!")
}
