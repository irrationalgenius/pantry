package bookRepository

import (
	"database/sql"
	"udemy-bookstore/models"
)

//BookRepository : BookRepository
type BookRepository struct{}

//GetBooks : GetBooks
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	// Invoke the db connection to get all book rows and assign
	// the collection values to "rows" and error to "err"
	rows, err := db.Query("select * from books")

	if err != nil {
		return []models.Book{}, err
	}

	// Close the db connection AFTER everything inside this handler
	// is executed.
	// defer rows.Close()

	// For each row in the data set retreieved from the database
	// set the hex value of the book variable
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	// If an error arises exit the execution with an empty Book slice and
	// error value
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

//GetBook : GetBook
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) (models.Book, error) {
	// Search the database for this parameter value
	row := db.QueryRow("select * from books where id=$1", id)
	// When value is located assign to the book pointer memory reference variable
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return book, err
}

//AddBook : AddBook
func (b BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {
	// Insert values received from the client into the database
	// Return the id of the insert into the hex value location for bookID
	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	// If an error arises exit the execution and return the error value
	if err != nil {
		return 0, err
	}

	return book.ID, nil
}

//UpdateBook : UpdateBook
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	// Update values received from the client to the record in the database
	// Return the id of the update into the hex value location for bookID
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	// If an error arises exit the execution and return the error value
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	// If an error arises exit the execution and return the error value
	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil
}

//RemoveBook : RemoveBook
func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	// Remove from the database a record matching this parameter value
	result, err := db.Exec("delete from books where id = $1", id)

	// If an error arises exit the execution and return the error value
	if err != nil {
		return 0, err
	}

	// Get the number of rows affected for the delete clause
	rowsDeleted, err := result.RowsAffected()

	// If an error arises exit the execution and return the error value
	if err != nil {
		return 0, err
	}

	return rowsDeleted, nil
}
