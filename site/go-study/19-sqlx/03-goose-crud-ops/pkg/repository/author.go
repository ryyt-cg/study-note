package repository

import "database/sql"

type AuthorEntity struct {
	ID    int64          `db:"id"`
	Name  string         `db:"name"`
	Bio   sql.NullString `db:"bio"`
	Books []BookEntity
}
