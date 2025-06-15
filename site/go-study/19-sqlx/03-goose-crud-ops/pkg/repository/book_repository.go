package repository

import (
	"database/sql"
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
	rows, err := repo.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

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
	row := repo.db.QueryRow("SELECT * FROM books WHERE id = ?", id)

	var book BookEntity
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// Insert a book
func (repo *BookRepository) Insert(book BookEntity) (*BookEntity, error) {
	result, err := repo.db.Exec("INSERT INTO books (title, author_id) VALUES (?, ?)", book.Title, book.AuthorID)
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
	_, err := repo.db.Exec("UPDATE books SET title = ?, author_id = ? WHERE id = ?", book.Title, book.AuthorID, book.ID)
	if err != nil {
		return err
	}

	return nil
}
