package models

import (
	"gorm.io/gorm"
)

// Attractions represents the attractions model
type Attractions struct {
	gorm.Model
	CustomID         uint               `gorm:"column:id" json:"id"`
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	OperatingHours   []OperatingHours   `gorm:"foreignKey:AttractionID" json:"operating_hours"`
	Content          string             `json:"content"`
	Images           StringArray        `gorm:"type:text[]" json:"images"`
	Location         Location           `gorm:"foreignKey:LocationID" json:"location"`
	LocationID       uint               `json:"location_id"`
	Price            int                `json:"price"`
	Rating           int                `json:"rating"`
	SocialMediaStack []SocialMediaStack `gorm:"foreignKey:AttractionID" json:"social_media_stack"`
	Details			 []Detail           `gorm:"foreignKey:AttractionID" json:"details"`
}
