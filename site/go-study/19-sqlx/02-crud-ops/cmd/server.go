package main

import (
	"02-crud-ops/pkg/db"
	"02-crud-ops/pkg/repository"
	"context"
	"database/sql"
	"log"
)

func main() {
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

	// Find all authors
	authors, err := authorRepository.FindAll()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("authors: %v\n", authors)

	// Find author by id
	author, err := authorRepository.FindById(1)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("author: %v\n", author)

	// find author with books by id
	authors, err = authorRepository.FindWithBooksById(6)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("author with books: %v\n", authors)

	// Insert an author
	newAuthor := repository.AuthorEntity{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	}
	insertedAuthor, err := authorRepository.Insert(newAuthor)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	log.Printf("inserted author: %v\n", insertedAuthor)
}
