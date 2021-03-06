CREATE TABLE users
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE authors
(
    id serial NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE poems
(
    id serial NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    text text NOT NULL
);

CREATE TABLE authors_list
(
    id serial NOT NULL UNIQUE,
    author_id int DEFAULT 0 REFERENCES authors (id) on delete CASCADE NOT NULL,
    poem_id int DEFAULT 0 REFERENCES poems (id) on delete CASCADE NOT NULL 
); 