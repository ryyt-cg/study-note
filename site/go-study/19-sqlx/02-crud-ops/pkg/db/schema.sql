CREATE TABLE authors (
    id   integer PRIMARY KEY AUTOINCREMENT,
    name text   NOT NULL,
    bio  text
);

CREATE TABLE books (
    id         integer PRIMARY KEY AUTOINCREMENT,
    title      text    NOT NULL,
    author_id  integer NOT NULL,
    published  date    NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors (id)
);

INSERT INTO authors (name, bio) VALUES ('William Shakespeare', 'English playwright and poet');
INSERT INTO authors (name, bio) VALUES ('Leo Tolstoy', 'Russian novelist');
INSERT INTO authors (name, bio) VALUES ('Fyodor Dostoevsky', 'Russian novelist');
INSERT INTO authors (name, bio) VALUES ('Gabriel Garcia Marquez', 'Colombian novelist');
INSERT INTO authors (name, bio) VALUES ('George Orwell', 'English novelist');
INSERT INTO authors (name, bio) VALUES ('J.R.R. Tolkien', 'English writer, poet, philologist, and university professor');

INSERT INTO books (title, author_id, published) VALUES ('Hamlet', 1, '1600-01-01');
INSERT INTO books (title, author_id, published) VALUES ('War and Peace', 2, '1869-01-01');
INSERT INTO books (title, author_id, published) VALUES ('Crime and Punishment', 3, '1866-01-01');
INSERT INTO books (title, author_id, published) VALUES ('One Hundred Years of Solitude', 4, '1967-01-01');
INSERT INTO books (title, author_id, published) VALUES ('Nineteen Eighty-Four', 5, '1949-01-01');
INSERT INTO books (title, author_id, published) VALUES ('The Lord of the Rings', 6, '1954-01-01');
INSERT INTO books (title, author_id, published) VALUES ('The Hobbit', 6, '1937-01-01');
INSERT INTO books (title, author_id, published) VALUES ('The Silmarillion', 6, '1977-01-01');
INSERT INTO books (title, author_id, published) VALUES ('The Children of Hurin', 6, '2007-01-01');
INSERT INTO books (title, author_id, published) VALUES ('The Fall of Gondolin', 6, '2018-01-01');
