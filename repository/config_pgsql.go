package repository

import (
	"database/sql"
	"errors"
	"os"
	"pantry/models"
)

// GetAppVars function in the repository is the implementation code
// behind the GetAppVars controller function.
func GetAppVars(db *sql.DB) error {

	var config models.ConfigParam

	// sqlConfigGet := `select prop_name, prop_category, prop_value from pantry.properties`
	sqlConfigGet := `select prop_name, prop_value from pantry.properties`

	rows, err := db.Query(sqlConfigGet)

	if err != nil {
		errorMsg := `Error: Issue occured while retrieving Config
      params from the database.`
		return errors.New(errorMsg)
	}

	for rows.Next() {
		// err = rows.Scan(&config.Name, &config.Category, &config.Value)
		err = rows.Scan(&config.Name, &config.Value)

		os.Setenv(config.Name, config.Value)
	}

	if err != nil {
		errorMsg := `Error: Issue occured while setting Config params`
		return errors.New(errorMsg)
	}

	return nil
}
