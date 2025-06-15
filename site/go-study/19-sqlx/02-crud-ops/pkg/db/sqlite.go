package db

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func SqliteConnect(ctx context.Context) (*sql.DB, error) {
	//:memory: is a special name that creates a new database that resides in memory rather than on disk.
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// read schema.sql file
	dat, err := os.ReadFile("pkg/db/schema.sql")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// convert []byte to string
	ddl := string(dat)
	// create tables and records
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}
