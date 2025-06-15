CREATE TABLE authors (
  id   integer PRIMARY KEY AUTOINCREMENT,
  name text    NOT NULL,
  bio  text
);

INSERT INTO authors (name, bio) VALUES ('William Shakespeare', 'English playwright and poet');
INSERT INTO authors (name, bio) VALUES ('Leo Tolstoy', 'Russian novelist');
INSERT INTO authors (name, bio) VALUES ('Fyodor Dostoevsky', 'Russian novelist');
INSERT INTO authors (name, bio) VALUES ('Gabriel Garcia Marquez', 'Colombian novelist');
INSERT INTO authors (name, bio) VALUES ('George Orwell', 'English novelist');
INSERT INTO authors (name, bio) VALUES ('J.R.R. Tolkien', 'English writer, poet, philologist, and university professor');
