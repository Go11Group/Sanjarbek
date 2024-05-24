CREATE TABLE authors (
    author_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);


CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    pages INT,
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES authors(author_id) ON DELETE SET NULL
);

