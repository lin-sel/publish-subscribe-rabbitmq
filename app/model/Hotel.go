package model

import uuid "github.com/satori/go.uuid"

// Hotel Contain ID, Name, Country, Address, X & Y Co-ordinates, Contact
type Hotel struct {
	ID          uuid.UUID   `json:"id" gorm:"type:varchar(36)"`
	HotelID     string      `json:"hotel_id" gorm:"type:varchar(20);unique"`
	Name        string      `json:"name" gorm:"type:varchar(30)"`
	Country     string      `json:"country" gorm:"type:varchar(10)"`
	Address     string      `json:"address" gorm:"type:varchar(50)"`
	Latitude    float64     `json:"latitude" gorm:"type:decimal"`
	Longitude   float64     `json:"longitude" gorm:"type:decimal"`
	Telephone   string      `json:"telephone" gorm:"type:varchar(15)"`
	Amenities   interface{} `json:"amenities" gorm:"type:json"`
	Description string      `json:"description" gorm:"type:varchar(1000)"`
	RoomCount   int         `json:"room_count" gorm:"type:int"`
	Currency    string      `json:"currency" gorm:"type:varchar(10)"`
}
