package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect() {
	var err error

	creds := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PSWD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := pgx.Connect(context.Background(), creds)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	taskTable := `CREATE TABLE IF NOT EXISTS tasks(
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			description TEXT,
			status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
			created_at TIMESTAMP DEFAULT now(),
			updated_at TIMESTAMP DEFAULT now())`

	_, err = db.Exec(context.Background(), taskTable)
	if err != nil {
		log.Fatal("Ошибка при создании БД:", err)
	}

	DB = db
	fmt.Println("Таблица успешно создана")
}
