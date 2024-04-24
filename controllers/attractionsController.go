package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/survani/pawadventures-backend/initializers"
	"github.com/survani/pawadventures-backend/models"
)

func AttractionsCreate(c *gin.Context) {
	var body struct {
		Name             string                    `json:"name" binding:"required"`
		Description      string                    `json:"description" binding:"required"`
		Content          string                    `json:"content" binding:"required"`
		OperatingHours   []models.OperatingHours   `json:"operating_hours"`
		Images           []string                  `json:"images" binding:"required"`
		LocationID       uint                      `json:"location_id" binding:"required"`
		Price            int                       `json:"price" binding:"required"`
		Rating           int                       `json:"rating" binding:"required"`
		SocialMediaStack []models.SocialMediaStack `json:"social_media_stack"`
		Details          []models.Detail           `json:"details"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert images to StringArray
	imagesArray := models.StringArray(body.Images)

	// create a house - this will tell the request what body it needs to update & where
	attraction := models.Attractions{
		Name:             body.Name,
		Description:      body.Description,
		OperatingHours:   body.OperatingHours,
		Content:          body.Content,
		Images:           imagesArray,
		LocationID:       body.LocationID,
		Price:            body.Price,
		Rating:           body.Rating,
		SocialMediaStack: body.SocialMediaStack,
		Details:          body.Details,
	}

	result := initializers.DB.Create(&attraction)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create attraction", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"attraction": attraction,
	})
}

func AttractionsIndex(c *gin.Context) {
	var attractions []models.Attractions
	initializers.DB.Preload("Location").Preload("OperatingHours").Preload("SocialMediaStack").Preload("Details").Find(&attractions)

	c.JSON(http.StatusOK, gin.H{
		"attractions": attractions,
	})
}

func AttractionsShowByID(c *gin.Context) {
	id := c.Param("id")

	var attraction models.Attractions
	result := initializers.DB.Preload("Location").Preload("OperatingHours").Preload("SocialMediaStack").Preload("Details").First(&attraction, id)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attraction not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attraction", "details": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attraction": attraction,
	})
}

func AttractionsShowBySlug(c *gin.Context) {
	slug := c.Param("slug")

	var attraction models.Attractions
	result := initializers.DB.Where("slug = ?", slug).Preload("Location").Preload("OperatingHours").Preload("SocialMediaStack").Preload("Details").First(&attraction)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attraction not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attraction", "details": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attraction": attraction,
	})
}

func AttractionsUpdate(c *gin.Context) {
	// Find The ID
	id := c.Param("id")

	var body struct {
		Name             string                    `json:"name" binding:"required"`
		Description      string                    `json:"description" binding:"required"`
		OperatingHours   []models.OperatingHours   `json:"operating_hours"`
		Content          string                    `json:"content" binding:"required"`
		Images           []string                  `json:"images" binding:"required"`
		LocationID       uint                      `json:"location_id" binding:"required"`
		Price            int                       `json:"price" binding:"required"`
		Rating           int                       `json:"rating" binding:"required"`
		SocialMediaStack []models.SocialMediaStack `json:"social_media_stack"`
		Details          []models.Detail           `json:"details"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the attraction we're updating
	var attraction models.Attractions
	result := initializers.DB.First(&attraction, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attraction not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attraction", "details": result.Error.Error()})
		}
		return
	}

	// Update the Attractions record
	result = initializers.DB.Model(&attraction).Updates(models.Attractions{
		Name:             body.Name,
		Description:      body.Description,
		OperatingHours:   body.OperatingHours,
		Content:          body.Content,
		Images:           body.Images,
		LocationID:       body.LocationID,
		Price:            body.Price,
		Rating:           body.Rating,
		SocialMediaStack: body.SocialMediaStack,
		Details:          body.Details,
	})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attraction", "details": result.Error.Error()})
		return
	}

	// Fetch the updated attraction record
	result = initializers.DB.Preload("Location").Preload("OperatingHours").Preload("SocialMediaStack").Preload("Details").First(&attraction, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated attraction", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"attraction": attraction,
	})
}

func AttractionsDelete(c *gin.Context) {
	id := c.Param("id")

	var attraction models.Attractions
	result := initializers.DB.First(&attraction, id)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Attraction not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve attraction", "details": result.Error.Error()})
		}
		return
	}

	result = initializers.DB.Delete(&attraction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete attraction", "details": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}