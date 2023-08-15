package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Init DB接続処理
func Init() (*sql.DB, error) {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		panic(err)
	}

	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, dbErr := sql.Open("postgres", dsn)
	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}
