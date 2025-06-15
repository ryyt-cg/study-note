package repository

import "time"

type BookEntity struct {
	ID          int64     `db:"id"`
	Title       string    `db:"title"`
	PublishedAt time.Time `db:"published_at"`
	AuthorID    int64     `db:"author_id"`
	Author      *AuthorEntity
}
