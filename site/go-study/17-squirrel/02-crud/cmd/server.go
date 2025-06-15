package main

import (
	"02-crud/pkg/db"
	"02-crud/pkg/repository"
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect to an in-memory database
	dbConn, err := db.SqliteConnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	authorRepository := repository.NewAuthorRepository(dbConn)

	// Insert an author
	author := repository.AuthorEntity{
		Name: "Brian Kernighan",
		Bio: sql.NullString{
			String: "Brian Wilson Kernighan is a Canadian computer scientist. He is a professor at the Department of Computer Science at Princeton University, USA.",
			Valid:  true,
		},
	}
	insertedAuthor, err := authorRepository.Insert(author)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("inserted author: %v\n", insertedAuthor)

	// List all authors
	authors, err := authorRepository.FindAll()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("authors: %v\n", authors)
}
