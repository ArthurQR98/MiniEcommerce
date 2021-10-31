package main

import (
	"log"

	"github.com/ArthurQR98/e-commerce/config"
	"github.com/ArthurQR98/e-commerce/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	if config.CheckConnection() == 0 {
		log.Fatal("Connection failed")
		return
	}
	routes.Routes()
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
