package db

import (
	"context"
	"database/sql"
	"log"
	"os"
)

func SqliteConnect(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	dat, err := os.ReadFile("migrations/00001_create_users.sql")
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
