package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

// Attractions represents the attractions model
type Attractions struct {
	gorm.Model
	CustomID         uint               `gorm:"column:id" json:"id"`
	Name             string             `json:"name"`
	Slug			 string             `gorm:"column:slug;unique_index" json:"slug"`
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

func (a *Attractions) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Slug == "" {
		a.Slug = slug.Make(a.Name)
	}
	return
}
