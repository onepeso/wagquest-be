package controllers

import (
	"fmt"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
	"github.com/survani/pawadventures-backend/initializers"
	"github.com/survani/pawadventures-backend/models"
)

func geocodeAddress(address string) (float64, float64, error) {
    geocoder := openstreetmap.Geocoder()
    location, err := geocoder.Geocode(address)
    if err != nil {
        return 0, 0, err
    }
    if location == nil {
        return 0, 0, fmt.Errorf("geocoding failed for address: %s", address)
    }
    return location.Lat, location.Lng, nil
}

func LocationCreate(c *gin.Context) {
    var body struct {
		Street  string `json:"street" binding:"required"`
        City    string `json:"city" binding:"required"`
        State   string `json:"state" binding:"required"`
        Zipcode string `json:"zipcode" binding:"required"`
        Country string `json:"country" binding:"required"`
    }

    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fullAddress := fmt.Sprintf("%s, %s, %s, %s, %s", body.Street, body.City, body.State, body.Zipcode, body.Country)
    latitude, longitude, err := geocodeAddress(fullAddress)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to geocode address"})
        return
    }

    location := models.Location{
		Street:   body.Street,
        City:      body.City,
        State:     body.State,
        Zipcode:   body.Zipcode,
        Country:   body.Country,
        Latitude:  latitude,
        Longitude: longitude,
    }

    result := initializers.DB.Create(&location)
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "location": location,
    })
}

func LocationIndex(c *gin.Context) {
    var locations []models.Location
    initializers.DB.Find(&locations)
    c.JSON(200, gin.H{
        "locations": locations,
    })
}

func LocationShow(c *gin.Context) {
    id := c.Param("id")
    var location models.Location
    initializers.DB.First(&location, id)
    c.JSON(200, gin.H{
        "location": location,
    })
}

func LocationUpdate(c *gin.Context) {
    id := c.Param("id")
    var body struct {
		Street  string `json:"street" binding:"required"`
        City    string `json:"city" binding:"required"`
        State   string `json:"state" binding:"required"`
        Zipcode string `json:"zipcode" binding:"required"`
        Country string `json:"country" binding:"required"`
    }

    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var location models.Location
    initializers.DB.First(&location, id)

    fullAddress := fmt.Sprintf("%s, %s, %s, %s", body.City, body.State, body.Zipcode, body.Country)
    latitude, longitude, err := geocodeAddress(fullAddress)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to geocode address"})
        return
    }

	location.Street = body.Street
    location.City = body.City
    location.State = body.State
    location.Zipcode = body.Zipcode
    location.Country = body.Country
    location.Latitude = latitude
    location.Longitude = longitude

    initializers.DB.Save(&location)

    c.JSON(200, gin.H{
        "location": location,
    })
}

func LocationDelete(c *gin.Context) {
    id := c.Param("id")
    initializers.DB.Delete(&models.Location{}, id)
    c.Status(200)
}