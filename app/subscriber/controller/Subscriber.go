package controller

import (
	"encoding/json"
	"fmt"

	"github.com/lin-sel/pub-sub-rmq/app/subscriber/model"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/service"
)

// SubscriberController have service
type SubscriberController struct {
	HotelService    *service.HotelService
	RatePlanService *service.RatePlanService
	RoomService     *service.RoomService
}

// NewSubscriberController Return New Object Of SubscriberController
func NewSubscriberController(hotelser *service.HotelService,
	rateplanser *service.RatePlanService, roomser *service.RoomService) *SubscriberController {
	return &SubscriberController{
		HotelService:    hotelser,
		RatePlanService: rateplanser,
		RoomService:     roomser,
	}
}

// Add Data to DB
func (sub *SubscriberController) Add(data []byte) {
	var hotel model.Hotel
	err := json.Unmarshal(data, &hotel)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = sub.HotelService.AddHotel(&hotel)
	if err != nil {
		fmt.Println(err.Error())
	}

	var room model.Room
	err = json.Unmarshal(data, &room)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = sub.RoomService.AddRoom(&room)
	if err != nil {
		fmt.Println(err.Error())
	}

	var ratePlan model.RatePlan
	err = json.Unmarshal(data, &ratePlan)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = sub.RatePlanService.AddRatePlan(&ratePlan)
	if err != nil {
		fmt.Println(err.Error())
	}

}
