package db

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteConnect(ctx context.Context) (*sql.DB, error) {
	//:memory: is a special name that creates a new database that resides in memory rather than on disk.
	// db, err := sql.Open("sqlite3", ":memory:")

	db, err := sql.Open("sqlite3", "goose-crud.sqlite")
	if err != nil {
		return nil, err
	}

	return db, nil
}
