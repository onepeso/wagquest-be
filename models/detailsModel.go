package models

import "gorm.io/gorm"

type Detail struct {
    gorm.Model
    AttractionID uint   `json:"attraction_id"`
    Name               string `json:"name"`
    Description        string `json:"description"`
}