package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect () *sql.DB {
	DB_HOST     := os.Getenv("DB_HOST");
	DB_PORT     := os.Getenv("DB_PORT");
	DB_USER     := os.Getenv("DB_USER");
	DB_PASSWORD := os.Getenv("DB_PASSWORD");
	DB_DATABASE := os.Getenv("DB_DATABASE");

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE);

	DB, err := sql.Open("postgres", psqlInfo);

	if err != nil {
		log.Fatalf("Error on try connect to db: %s", err)
	};

	return DB;
}