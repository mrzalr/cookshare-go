package api

import (
	"log"

	"github.com/joho/godotenv"
)

func StartApplication() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error when load .env files | err : %v", err.Error())
	}
}
