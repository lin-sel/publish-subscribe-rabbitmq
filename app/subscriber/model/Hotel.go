package model

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// Hotel Contain ID, Name, Country, Address, X & Y Co-ordinates, Contact
type Hotel struct {
	ID          uuid.UUID       `json:"id" gorm:"type:varchar(36)"`
	HotelID     string          `json:"hotel_id" gorm:"type:varchar(30);unique"`
	Name        string          `json:"name" gorm:"type:text"`
	Country     string          `json:"country" gorm:"type:text"`
	Address     string          `json:"address" gorm:"type:text"`
	Latitude    float64         `json:"latitude" gorm:"type:decimal"`
	Longitude   float64         `json:"longitude" gorm:"type:decimal"`
	Telephone   string          `json:"telephone" gorm:"type:text"`
	Amenities   json.RawMessage `json:"amenities" gorm:"type:json"`
	Description string          `json:"description" gorm:"type:text"`
	RoomCount   int             `json:"room_count" gorm:"type:int"`
	Currency    string          `json:"currency" gorm:"type:text"`
}
