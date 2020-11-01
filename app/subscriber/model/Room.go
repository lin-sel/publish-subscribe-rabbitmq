package model

import uuid "github.com/satori/go.uuid"

// Room Contain ID, HotelID, Description, Name, Capacity
type Room struct {
	ID          uuid.UUID   `json:"id" gorm:"type:varchar(36)"`
	RoomID      string      `json:"room_id" gorm:"type:varchar(20);unique"`
	HotelID     string      `json:"hotel_id" gorm:"type:varchar(20)" gorm:"foreignKey:hotel_id;references:hotels(id)"`
	Description string      `json:"description" gorm:"type:varchar(1000)"`
	Name        string      `json:"name" gorm:"type:varchar(20)"`
	Capacity    interface{} `json:"capacity" gorm:"type:json"`
}
