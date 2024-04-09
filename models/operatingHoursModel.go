package models

import "gorm.io/gorm"

type OperatingHours struct {
	gorm.Model
	AttractionID uint
	Day          string `json:"day"`
	Hours        string `json:"hours"`
}
