package main

import (
	"log"

	"github.com/BenjaminLai/kris-kringle-generator/app"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Generate()
}
