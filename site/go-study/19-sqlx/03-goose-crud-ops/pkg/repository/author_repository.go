package repository

import (
	"database/sql"
	"log"
)

type AuthorRepositorier interface {
	FindAll() ([]AuthorEntity, error)
	FindById(id int) (*AuthorEntity, error)
	FindWithBooksById(id int) ([]AuthorEntity, error)
	Insert(authorEntity AuthorEntity) (*AuthorEntity, error)
	Update(authorEntity AuthorEntity) (*AuthorEntity, error)
}

// AuthorRepository - CRUD operations for authors
type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func rowsClose(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
}

// FindAll
// find all authors
func (repo *AuthorRepository) FindAll() ([]AuthorEntity, error) {
	rows, err := repo.db.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	defer rowsClose(rows)

	var authors []AuthorEntity
	for rows.Next() {
		var author AuthorEntity
		err := rows.Scan(&author.ID, &author.Name, &author.Bio)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}

// FindById
// find the author by id
func (repo *AuthorRepository) FindById(id int) (*AuthorEntity, error) {
	row := repo.db.QueryRow("SELECT * FROM authors WHERE id = ?", id)

	var author AuthorEntity
	err := row.Scan(&author.ID, &author.Name, &author.Bio)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

// FindWithBooksById
// find the author by id with books
func (repo *AuthorRepository) FindWithBooksById(id int) ([]AuthorEntity, error) {
	query := `SELECT a.id, a.name, a.bio, b.id, b.title, b.author_id FROM authors a
					LEFT JOIN books b ON a.id = b.author_id
					WHERE a.id = ?`
	rows, err := repo.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rowsClose(rows)

	var authors []AuthorEntity
	for rows.Next() {
		var author AuthorEntity
		var book BookEntity
		err := rows.Scan(&author.ID, &author.Name, &author.Bio, &book.ID, &book.Title, &book.AuthorID)
		if err != nil {
			return nil, err
		}
		author.Books = append(author.Books, book)
		authors = append(authors, author)
	}

	return authors, nil
}

// Insert an author
func (repo *AuthorRepository) Insert(author AuthorEntity) (*AuthorEntity, error) {
	result, err := repo.db.Exec("INSERT INTO authors (name, bio) VALUES (?, ?)", author.Name, author.Bio)
	if err != nil {
		return nil, err
	}

	author.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &author, nil
}

// Update an author
func (repo *AuthorRepository) Update(author AuthorEntity) (*AuthorEntity, error) {
	_, err := repo.db.Exec("UPDATE authors SET name = ?, bio = ? WHERE id = ?", author.Name, author.Bio, author.ID)
	if err != nil {
		return nil, err
	}
	return &author, nil
}
