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

	// See logFatal() function
	// logFatal(err)
	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}
