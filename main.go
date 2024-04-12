// NOTE: Make sure we add our update to github
// TODO: Change the flag images to be non-copyrighted ones.

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/survani/pawadventures-backend/controllers"
	"github.com/survani/pawadventures-backend/initializers"
)

func init() {
	// LoadEnvVariables()
	initializers.LoadEnvVariables()
	// ConnectToDB()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000",
    "http://10.0.0.124:3000",
    "https://wagquest-fe.vercel.app",
    "https://wagquest-fe-git-main-onepeso-studio.vercel.app",
    "https://wagquest-3ifydryp0-onepeso-studio.vercel.app",
    "https://wagquest.com",} // Update this when deploying
	router.Use(cors.New(config))

	// ATTRACTION API ENDPOINT ROUTES
	router.GET("/attractions", controllers.AttractionsIndex)
	router.GET("/attraction/:id", controllers.AttractionsShowByID)
	router.GET("/attraction/slug/:slug", controllers.AttractionsShowBySlug)
	router.POST("/attraction", controllers.AttractionsCreate)
	router.PUT("/attraction/:id", controllers.AttractionsUpdate)
	router.DELETE("/attraction/:id", controllers.AttractionsDelete)

	// LOCATION API ENDPOINT ROUTES
	router.GET("/locations", controllers.LocationIndex)
	router.GET("/location/:id", controllers.LocationShow)
	router.POST("/locations", controllers.LocationCreate)
	router.PUT("/location/:id", controllers.LocationUpdate)
	router.DELETE("/location/:id", controllers.LocationDelete)

	router.Run()
}
