package controllers

import (
	"database/sql"
	"fmt"
	"pantry/models"
	"pantry/repository"
	"pantry/utils"
)

// GetAppVars is for retreieving application parameters
// from the database.
func GetAppVars(db *sql.DB) []models.ConfigParam {

	configsRepo := repository.ConfigsRepository{}

	configParams, err := configsRepo.GetAppVars(db)

	if err != nil {
		utils.LogFatal(err)
		return []models.ConfigParam{}
	}

	return configParams
}

// SplashScreen is for the application startup, which will
// output application parameters.
func SplashScreen(configParams []models.ConfigParam) {

	var splashScreenMsg string

	for _, configParam := range configParams {
		if configParam.Name == "APP_NAME" {

			splashScreenMsg = "Starting up the"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
		if configParam.Name == "APP_VERSION" {

			splashScreenMsg = "Application Version:"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
		if configParam.Name == "APP_PURPOSE" {

			splashScreenMsg = "Purpose:"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
		if configParam.Name == "APP_OWNER" {

			splashScreenMsg = "Owner:"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
		if configParam.Name == "APP_DEVELOPER" {

			splashScreenMsg = "Developer:"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
		if configParam.Name == "APP_DEVELOPER_EMAIL" {

			splashScreenMsg = "Developer Email:"
			fmt.Println(splashScreenMsg, configParam.Value)
		}
	}
}
