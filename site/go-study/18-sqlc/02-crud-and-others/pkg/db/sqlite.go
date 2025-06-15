package db

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func SqliteConnect(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile("pkg/db/schema.sql")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ddl := string(dat)
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}
