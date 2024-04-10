package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func New() {
	newPostgresDB()
}

func newPostgresDB() {

	postgres_uri := os.Getenv("POSTGRES_URI")

	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", postgres_uri)
		if err != nil {
			log.Fatalf("sql.Open(): %v", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalf("db.Ping(): %v", err)
		}

		fmt.Println("New connection established to postgres")
	})
}

func Pool() *sql.DB {
	return db
}
