package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/pub-sub-rmq/app"
	"github.com/lin-sel/pub-sub-rmq/app/log"
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
	app := app.NewApp(db, log)
	app.TableMigration()
}
