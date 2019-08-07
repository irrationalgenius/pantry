package repository

import (
	"database/sql"
	"errors"
	"pantry/models"
)

//ConfigsRepository : ConfigsRepository
type ConfigsRepository struct{}

// GetAppVars function in the repository is the implementation code
// behind the GetAppVars controller function.
func (c ConfigsRepository) GetAppVars(db *sql.DB) ([]models.ConfigParam, error) {

	var config models.ConfigParam
	var configs []models.ConfigParam

	sqlConfigGet := `select prop_name, prop_category, prop_value from pantry.properties`

	rows, err := db.Query(sqlConfigGet)

	if err != nil {
		errorMsg := `Error: Issue occured while retrieving Config
      params from the database.`
		return []models.ConfigParam{}, errors.New(errorMsg)
	}

	for rows.Next() {
		err = rows.Scan(&config.Name, &config.Category, &config.Value)

		configs = append(configs, config)
	}

	if err != nil {
		errorMsg := `Error: Issue occured while setting Config
      params`
		return []models.ConfigParam{}, errors.New(errorMsg)
	}

	return configs, nil
}
