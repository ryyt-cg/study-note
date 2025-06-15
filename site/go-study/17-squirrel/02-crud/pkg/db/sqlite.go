package db

import (
	"context"
	"database/sql"
	"log"
	"os"
)

// SqliteConnect connects to an in-memory database
// and creates tables from schema.sql
func SqliteConnect(ctx context.Context) (*sql.DB, error) {
	// connect to an in-memory database
	db, err := sql.Open("sqlite3", "squirrel-crud.db")
	if err != nil {
		return nil, err
	}

	// read schema.sql
	dat, err := os.ReadFile("migrations/schema.sql")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ddl := string(dat)
	// create tables from schema.sql
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	return db, nil
}
