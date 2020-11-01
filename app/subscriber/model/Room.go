package model

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

// Room Contain ID, HotelID, Description, Name, Capacity
type Room struct {
	ID          uuid.UUID       `json:"id" gorm:"type:varchar(36)"`
	RoomID      string          `json:"room_id" gorm:"type:varchar(30);unique"`
	HotelID     string          `json:"hotel_id" gorm:"type:varchar(30)" gorm:"foreignKey:hotel_id;references:hotels(id)"`
	Description string          `json:"description" gorm:"type:text"`
	Name        string          `json:"name" gorm:"type:text"`
	Capacity    json.RawMessage `json:"capacity" gorm:"type:json"`
}
