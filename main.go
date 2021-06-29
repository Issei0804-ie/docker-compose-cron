package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("/var/log/hoge.log", os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error())
	}
	logger := log.Logger{}
	logger.SetOutput(f)

	logger.Println("Start program")

	MODE := os.Getenv("MODE")

	if MODE == "PRODUCTION" {
		logger.Println("MODE is PRODUCTION")
	} else if MODE == "DEVELOPMENT" {
		logger.Println("MODE is DEVELOPMENT")
	} else {
		logger.Println("Mode is not setting")
	}

	logger.Println("Finish program")
}
