package repository

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
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
	query, _, _ := sq.Select("id", "name", "bio").From("authors").ToSql()
	log.Printf("find all authors query: %v\n", query)
	rows, err := repo.db.Query(query)
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
	query, args, _ := sq.Select("id", "name", "bio").From("authors").Where(sq.Eq{"id": id}).ToSql()
	log.Printf("find author by id query: %v\n", query)
	log.Printf("find author by id args: %v", args)
	row := repo.db.QueryRow(query, args[0])
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
	query, args, _ := sq.Select("a.id", "a.name", "a.bio").
		From("authors a").
		Join("books b ON a.id = b.author_id").
		Where(sq.Eq{"a.id": id}).
		ToSql()

	log.Printf("find author with books by id query: %v\n", query)
	log.Printf("find author with books by id args: %v", args)
	rows, err := repo.db.Query(query, args[0])
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

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

// Insert an author
func (repo *AuthorRepository) Insert(author AuthorEntity) (*AuthorEntity, error) {
	query, args, _ := sq.Insert("authors").Columns("name", "bio").
		Values(author.Name, author.Bio).
		ToSql()

	log.Printf("insert author query: %v\n", query)
	log.Printf("insert author args: %v", args)

	result, err := repo.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	author.ID, err = result.LastInsertId()
	return &author, nil
}

// Update an author
func (repo *AuthorRepository) Update(authorEntity AuthorEntity) (*AuthorEntity, error) {
	query, args, _ := sq.Update("authors").
		Set("name", authorEntity.Name).
		Set("bio", authorEntity.Bio).
		Where(sq.Eq{"id": authorEntity.ID}).
		ToSql()

	log.Printf("update author query: %v\n", query)
	log.Printf("update author args: %v", args)
	_, err := repo.db.Exec(query, args[0], args[1], args[2])
	if err != nil {
		return nil, err
	}

	return &authorEntity, nil
}
