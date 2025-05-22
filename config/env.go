package config

import (
	"log"

	"github.com/joho/godotenv"
)

func InitConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка файла .env")
		panic(err)
	}
	log.Println("Конфиг .env загружен")
}
