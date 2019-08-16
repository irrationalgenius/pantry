package controllers

import (
	"database/sql"
	"fmt"
	"os"
	"pantry/repository"
	"pantry/utils"
)

// GetAppVars is for retreieving application parameters
// from the database.
func GetAppVars(db *sql.DB) {

	err := repository.GetAppVars(db)

	if err != nil {
		utils.LogFatal(err)
	}

	// Display the splashscreen through the terminal
	splashScreen()
}

// SplashScreen is for the application startup, which will
// output application parameters.
func splashScreen() {

	fmt.Println("Starting up the", os.Getenv("APP_NAME"))
	fmt.Println("Application Version:", os.Getenv("APP_VERSION"))
	fmt.Println("Purpose:", os.Getenv("APP_PURPOSE"))
	fmt.Println("Owner:", os.Getenv("APP_OWNER"))
	fmt.Println("Developer:", os.Getenv("APP_DEVELOPER"))
	fmt.Println("Developer Email:", os.Getenv("APP_DEVELOPER_EMAIL"))
	fmt.Println("Current Visit Wait Interval (in days):", os.Getenv("APP_VISIT_INTERVAL"))

}
