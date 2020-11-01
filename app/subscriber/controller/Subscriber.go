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
	var js model.JSON
	err := json.Unmarshal(data, &js)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, data := range js.Offers {
		hotel := data.Hotel
		err = sub.HotelService.AddHotel(&hotel)
		if err != nil {
			fmt.Println(err.Error())
		}

		room := data.Room
		err = sub.RoomService.AddRoom(&room)
		if err != nil {
			fmt.Println(err.Error())
		}

		ratePlan := data.RatePlan
		err = sub.RatePlanService.AddRatePlan(&ratePlan)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
