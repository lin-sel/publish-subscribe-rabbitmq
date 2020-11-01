package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/pub-sub-rmq/app"
	"github.com/lin-sel/pub-sub-rmq/app/log"
	"github.com/lin-sel/pub-sub-rmq/app/repository"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/controller"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/service"
)

func getConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", "swabhav:swabhav@/test?charset=utf8&parseTime=True&loc=Local")
	return
}

func main() {
	log := log.NewLogger()
	db, err := getConnection()
	if err != nil {
		log.Fatalf("unable to connecto db error:%s", err.Error())
	}
	db.LogMode(true)
	route := mux.NewRouter()
	app := app.NewApp(db, log, route)
	repo := repository.NewRepo()
	hotelSer := service.NewHotelService(db, repo)
	roomSer := service.NewRoomService(db, repo)
	ratePlanSer := service.NewRatePlanService(db, repo)
	subContr := controller.NewSubscriberController(hotelSer, ratePlanSer, roomSer)
	app.InitApp(subContr)
}
