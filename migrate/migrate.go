package main

import (
	"fmt"

	"github.com/survani/pawadventures-backend/initializers"
	"github.com/survani/pawadventures-backend/models"
)

func init() {
	// LoadEnvVariables()
	initializers.LoadEnvVariables()
	// ConnectToDB()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB = initializers.DB.Debug()
	fmt.Println("AutoMigrate executed")
	err := initializers.DB.AutoMigrate(&models.Attractions{}, &models.Location{}, &models.OperatingHours{}, &models.Detail{}, &models.SocialMediaStack{})
	if err != nil {
		fmt.Println("Error during AutoMigrate:", err)
	}
}
