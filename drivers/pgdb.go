package drivers

import (
	"database/sql"
	"os"
	"pantry/utils"

	"github.com/lib/pq"
)

var db *sql.DB

//PgConnect : PgConnect
func PgConnect() *sql.DB {
	// Get the database connection string by using the
	// systems environmental variable, by the Gotenv package
	pgHost, err := pq.ParseURL(os.Getenv("PG_URL"))

	// See logFatal() function
	utils.LogFatal(err)

	// Connect to the database, and return a handler and
	// err vars "db", "err"
	db, err = sql.Open("postgres", pgHost)

	// See logFatal() function
	utils.LogFatal(err)

	// Verifies an active db connection
	err = db.Ping()

	// See logFatal() function
	utils.LogFatal(err)

	return db
}
