package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

// Book : Book Struct
type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

// Declare a slice data type of the Book struct
var books []Book

/*
	Clarification of Pointers
		* Retrieves the data from memory location (Dereference)
		& Retreieves the hex value from the memory location
*/
var db *sql.DB

// Before the procedure does anything, this is the first function
// which will actuate, loading the systems environmental vars
func init() {
	gotenv.Load()

}

// If for any reason, something goes amiss then exit and
// send the message to the logging console.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Get the database connection string by using the
	// systems environmental variable, by the Gotenv package
	pgHost, err := pq.ParseURL(os.Getenv("PG_URL"))

	// See logFatal() function
	logFatal(err)

	// Connect to the database, and return a handler and
	// err vars "db", "err"
	db, err = sql.Open("postgres", pgHost)

	// See logFatal() function
	logFatal(err)

	// Verifies an active db connection
	err = db.Ping()

	// See logFatal() function
	logFatal(err)

	// Create an instance of the New Router function in mux
	router := mux.NewRouter()

	// Map each URL route to a speific handler function
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	// Run the server, if any errors exits then logFatal()
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Geting all books method is invoked")

	// Initialize a variable with the type of the Book struct
	var book Book

	// Assign an empty slice to the global []books data type
	books = []Book{}

	// Invoke the db connection to get all book rows and assign
	// the collection values to "rows" and error to "err"
	rows, err := db.Query("select * from books")
	logFatal(err)

	// Close the db connection AFTER everything inside this handler
	// is executed.
	defer rows.Close()

	// For each row in the data set retreieved from the database
	// set the hex value of the book variable
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

		// See logFatal() function
		logFatal(err)

		// Handle the books slice array by adding to the end of the array
		// after being Initialized as empty
		books = append(books, book)
	}

	// Convert the books slice array to a json object and send the response
	// to the client as "w"
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Retrieving a single book method is invoked")

	// Initialize a variable with the type of the Book struct
	var book Book

	// Retreieve the URL parameters as "r" and insert into a
	// map data type as "params"
	params := mux.Vars(r)

	// Search the database for this parameter value
	row := db.QueryRow("select * from books where id=$1", params["id"])

	// Insert the values from the database into the hex value
	// of the book vars
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	// See logFatal() function
	logFatal(err)

	// Convert the book data type to a json object and send the response
	// to the client as "w"
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adding a book method is invoked")

	// Initialize a variable with the type of the Book struct
	var book Book

	// Initialize a variable with the type of int
	var bookID int

	// Handle the response Body and map values to the hex value
	// of the book var
	json.NewDecoder(r.Body).Decode(&book)

	// Insert values received from the client into the database
	// Return the id of the insert into the hex value location for bookID
	err := db.QueryRow("insert into books(title, author, year) values($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year).Scan(&bookID)

	// See logFatal() function
	logFatal(err)

	// Convert the bookID var to a json object and send the response
	// to the client as "w"
	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating a book method is invoked")

	// Initialize a variable with the type of the Book struct
	var book Book

	// Retrieves the response body and maps it to the book variable
	json.NewDecoder(r.Body).Decode(&book)

	// Update values received from the client to the record in the database
	// Return the id of the update into the hex value location for bookID
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	// See logFatal() function
	logFatal(err)

	// Get the number of rows affected for the update clause
	rowUpdated, err := result.RowsAffected()

	// See logFatal() function
	logFatal(err)

	// Convert the rowUpdated var to a json object and send the response
	// to the client as "w"
	json.NewEncoder(w).Encode(rowUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removing a book method is invoked")

}
