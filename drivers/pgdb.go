package drivers

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

// If for any reason, something goes amiss then exit and
// send the message to the logging console.
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//PgConnect : PgConnect
func PgConnect() *sql.DB {
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

	return db
}
