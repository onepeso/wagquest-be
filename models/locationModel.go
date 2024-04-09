// TODO: Maybe we can add a way to add locations to the house by City or State or Zipcode or Country. Then return the house with the location.

package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	CustomID uint   `gorm:"column:id" json:"id"`
	Street  string `json:"street"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zipcode  string `json:"zipcode"`
	Country  string `json:"country"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
