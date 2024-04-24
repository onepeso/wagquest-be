package models

import (
	"time"
)

type OperatingHours struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	AttractionID uint      `gorm:"index" json:"attraction_id"`
	Day          string    `gorm:"not null" json:"day"`
	OpenTime     time.Time `gorm:"not null" json:"open_time"`
	CloseTime    time.Time `gorm:"not null" json:"close_time"`
}