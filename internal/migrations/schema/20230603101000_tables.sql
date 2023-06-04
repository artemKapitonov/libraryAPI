-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL
);

CREATE TABLE user_books (
    user_id INTEGER REFERENCES users(id),
    book_id INTEGER REFERENCES books(id),
    PRIMARY KEY (user_id, book_id)
);

-- +goose Down
DROP TABLE user_books;

DROP TABLE users;

DROP TABLE books;