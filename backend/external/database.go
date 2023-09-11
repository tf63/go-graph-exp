package external

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// GORMでDBに接続し，GORMのDBオブジェクトを返す
func ConnectDatabase() (*sql.DB, error) {

	// DBへの接続に必要な情報
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_NAME"),
		os.Getenv("POSTGRES_PORT"),
	)

	// DBへ接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
