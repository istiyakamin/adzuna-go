package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ADZUNA_APP_ID, ADZUNA_APP_KEY string

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ADZUNA_APP_ID = os.Getenv("ADZUNA_APP_ID")
	ADZUNA_APP_KEY = os.Getenv("ADZUNA_APP_KEY")
}