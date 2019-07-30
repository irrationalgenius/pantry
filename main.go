package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"udemy-bookstore/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

// Declare a slice data type of the Book struct
var books []models.Book
var db *sql.DB

/*
	Clarification of Pointers
		* Retrieves the value from the memory location (Dereference)
		& Retreieves the hex value from the memory location
*/

// Before the procedure does anything, this is the first function
// which will actuate, loading the systems environmental vars
func init() {
	gotenv.Load()
}

func main() {
	// Connect to the database, and verify the connection, if any issues
	// program exits immediately
	// db = drivers.PgConnect()

	// Create an instance of the Contoller object, and assign its value to booker
	// guest := controllers.Controller{}

	// controllers.GuestTest()

	// Create an instance of the New Router function in mux
	router := mux.NewRouter()

	// Map each URL route to a speific handler function
	// router.HandleFunc("/guests", controllers.GuestTest()).Methods("GET")
	// router.HandleFunc("/book/{id}", booker.GetBook(db)).Methods("GET")
	// router.HandleFunc("/book", booker.AddBook(db)).Methods("POST")
	// router.HandleFunc("/book", booker.UpdateBook(db)).Methods("PUT")
	// router.HandleFunc("/book/{id}", booker.RemoveBook(db)).Methods("DELETE")

	// Run the server, if any errors exits then logFatal()
	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
