-- Persist data in a database

CREATE TABLE books(
  id serial,
  title varchar,
  author varchar,
  year varchar
);

INSERT INTO books (title, author, year) VALUES ('Golang is Great', 'Mr. Great', '2012');
INSERT INTO books (title, author, year) VALUES ('C++ is the Greatest', 'Mr. C++', '2015');
-- INSERT INTO books (title, author, year) VALUES ('C++ is the Greatest', 'Mr. C++', '2015');
