-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: GetBook :one
SELECT * FROM books
WHERE id = ? LIMIT 1;

-- name: CreateBook :one
INSERT INTO books (
    title, published, author_id
) VALUES (?, ?, ?)
    RETURNING *;

-- name: UpdateBook :one
UPDATE books
set title = ?,
    published = ?,
    author_id = ?
WHERE id = ?
RETURNING *;