-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY unique,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL unique,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY unique,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    path TEXT NOT NULL
);

CREATE TABLE user_books (
    id      serial not null unique,
    user_id INTEGER REFERENCES users(id),
    book_id INTEGER REFERENCES books(id),
    PRIMARY KEY (user_id, book_id)
);

-- +goose Down
DROP TABLE user_books;

DROP TABLE users;

DROP TABLE books;