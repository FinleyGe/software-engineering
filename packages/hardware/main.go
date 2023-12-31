package main

import (
	"github.com/joho/godotenv"
	"hardware/cli"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cli.Run()
}
