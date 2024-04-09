package models

import "gorm.io/gorm"

type SocialMediaStack struct {
	gorm.Model
	AttractionID uint
	Platform     string `json:"platform"`
	Handle       string `json:"handle"`
}
