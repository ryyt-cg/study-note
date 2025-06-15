package main

import (
	"02-crud-and-others/pkg/db"
	"02-crud-and-others/pkg/db/sqlgen"
	"02-crud-and-others/pkg/repository"
	"context"
	"database/sql"
	"log"
	"reflect"
	"time"
)

func main() {
	ctx := context.Background()

	dbConn, err := db.SqliteConnect(ctx)
	if err != nil {
		log.Panic(err)
	}
	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Panic(err)
		}
	}(dbConn)

	queries := sqlgen.New(dbConn)
	authorRepository := repository.NewAuthorRepository(ctx, queries)

	// list all authors
	authors, err := authorRepository.FindAll()
	log.Println(authors)

	newAuthor := sqlgen.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	}

	// create an author
	insertedAuthor, err := authorRepository.Insert(newAuthor)
	if err != nil {
		log.Panic(err)
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := authorRepository.FindById(insertedAuthor.ID)
	if err != nil {
		log.Panic(err)
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))

	// delete the author we just inserted
	//err = authorRepository.Delete(insertedAuthor.ID)
	//if err != nil {
	//	log.Panic(err)
	//}

	// list all authors
	authors, err = authorRepository.FindAll()
	log.Println(authors)

	bookRepository := repository.NewBookRepository(ctx, queries)
	// create some book
	newBook := sqlgen.CreateBookParams{
		Title:     "The Go Programming Language",
		Published: time.Now(),
		AuthorID:  1,
	}
	insertedBook, _ := bookRepository.Insert(newBook)

	newBook = sqlgen.CreateBookParams{
		Title:     "The C Programming Language",
		Published: time.Now(),
		AuthorID:  1,
	}
	insertedBook, _ = bookRepository.Insert(newBook)

	log.Println(insertedBook)

	// list all books
	books, err := bookRepository.FindAll()
	log.Println(books)

	// list all authors with books
	authorsWithBooks, _ := authorRepository.FindWithBooksById(1)
	log.Println(authorsWithBooks)
}
