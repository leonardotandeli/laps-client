package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var StringBanco = ""

func CarregarConfiguracao() {

	erro := godotenv.Load()
	if erro != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	StringBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

}
