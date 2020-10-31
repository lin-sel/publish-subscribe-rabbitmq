package model

import uuid "github.com/satori/go.uuid"

// RatePlan Contain ID, HotelID, MealPlan
type RatePlan struct {
	ID                 uuid.UUID   `json:"id" gorm:"type:varchar(36)"`
	RatePlanID         string      `json:"rate_plan_id" gorm:"type:varchar(20);unique"`
	HotelID            string      `json:"hotel_id" gorm:"type:varchar(20)" gorm:"foreignKey:hotel_id;references:hotels(id)"`
	MealPlan           string      `json:"meal_plan" gorm:"type:varchar(50)"`
	CancellationPolicy interface{} `json:"cancellation_policy" gorm:"type:json"`
	OtherConditions    interface{} `json:"other_conditions" gorm:"type:json"`
}
