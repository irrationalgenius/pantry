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
	appVars = controllers.GetAppVars(db)

	// Displays the splashscreen by the console
	controllers.SplashScreen(appVars)

	// Create an instance of the Contoller object, and assign its value to booker
	guest := controllers.GuestController{}

	// Create an instance of the New Router function in mux
	router := mux.NewRouter()

	// Map each URL route to a speific handler function
	router.HandleFunc("/guests", guest.GetGuests(db)).Methods("GET")
	router.HandleFunc("/guest/{id}", guest.GetGuest(db)).Methods("GET")
	router.HandleFunc("/guest", guest.AddGuest(db)).Methods("POST")
	router.HandleFunc("/guest", guest.UpdateGuest(db)).Methods("PUT")
	router.HandleFunc("/guest/{id}", guest.RemoveGuest(db)).Methods("DELETE")

	// Run the server, if any errors exits then logFatal()
	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
