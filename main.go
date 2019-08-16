package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"pantry/controllers"
	"pantry/drivers"
	"pantry/models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sql.DB
var appVars []models.ConfigParam

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
	db = drivers.PgConnect()

	// Retrieves all application config params from the database
	controllers.GetAppVars(db)

	// Create an instance of the Contoller object, and assign its value to guest
	guest := controllers.GuestController{}
	visit := controllers.VisitController{}

	// Create an instance of the New Router function in mux
	router := mux.NewRouter()

	// Map each URL route to a speific handler function
	// Handle all Guest object states
	router.HandleFunc("/guests", guest.GetGuests(db)).Methods("GET")
	router.HandleFunc("/guests/{id}", guest.GetGuest(db)).Methods("GET")
	router.HandleFunc("/guests", guest.AddGuest(db)).Methods("POST")
	router.HandleFunc("/guests", guest.UpdateGuest(db)).Methods("PUT")
	router.HandleFunc("/guests/{id}", guest.UpdateGuest(db)).Methods("PUT")
	router.HandleFunc("/guests/{id}", guest.RemoveGuest(db)).Methods("DELETE")

	// Handle all Visit object states per Guest object
	router.HandleFunc("/guests/{id}/visits", visit.GetGuestVisits(db)).Methods("GET")
	router.HandleFunc("/guests/{id}/visits", visit.AddGuestVisit(db)).Methods("POST")

	// Run the server, if any errors exits then logFatal()
	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Server is running at port 8000")

}
