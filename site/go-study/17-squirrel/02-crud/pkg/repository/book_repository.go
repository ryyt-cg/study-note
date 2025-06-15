package repository

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type BookRepositorier interface {
	FindAll() ([]BookEntity, error)
	FindById(id int) (*AuthorEntity, error)
	Insert(book BookEntity) (*BookEntity, error)
	Update(book BookEntity) error
}

// BookRepository searches books from the database
type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

// FindAll
// find all books
func (repo *BookRepository) FindAll() ([]BookEntity, error) {
	query, _, _ := sq.Select("id", "title", "author_id").From("books").ToSql()
	log.Printf("find all books query: %v\n", query)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rowsClose(rows)

	var books []BookEntity
	for rows.Next() {
		var book BookEntity
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// FindById
// find the book by id
func (repo *BookRepository) FindById(id int) (*BookEntity, error) {
	query, args, _ := sq.Select("id", "title", "author_id").From("books").Where(sq.Eq{"id": id}).ToSql()
	log.Printf("find book by id query: %v\n", query)
	row := repo.db.QueryRow(query, args...)

	var book BookEntity
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Insert a book
func (repo *BookRepository) Insert(book BookEntity) (*BookEntity, error) {
	query, args, _ := sq.Insert("books").Columns("title", "author_id").Values(book.Title, book.AuthorID).ToSql()
	log.Printf("insert a book query: %v\n", query)
	log.Printf("insert a book args: %v\n", args)
	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	book.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Update a book
func (repo *BookRepository) Update(book BookEntity) error {
	query, args, _ := sq.Update("books").
		Set("title", book.Title).
		Set("published_at", book.PublishedAt).
		Where(sq.Eq{"id": book.ID}).ToSql()
	log.Printf("update a book query: %v\n", query)
	log.Printf("update a book args: %v\n", args)
	_, err := repo.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
