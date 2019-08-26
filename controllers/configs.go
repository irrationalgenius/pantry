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

	fmt.Println("Starting up the", os.Getenv("app_name"))
	fmt.Println("Application Version:", os.Getenv("app_version"))
	fmt.Println("Purpose:", os.Getenv("app_purpose"))
	fmt.Println("Owner:", os.Getenv("app_owner"))
	fmt.Println("Developer:", os.Getenv("app_developer"))
	fmt.Println("Developer Email:", os.Getenv("app_developer_email"))
	fmt.Println("Current Visit Wait Interval (in days):", os.Getenv("app_visit_interval"))

}
