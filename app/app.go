package app

import (
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/pub-sub-rmq/app/model"
	"github.com/sirupsen/logrus"
)

// App DB, Log
type App struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

// NewApp Return New Object Of App
func NewApp(db *gorm.DB, lg *logrus.Logger) *App {
	return &App{
		DB:  db,
		Log: lg,
	}
}

// TableMigration Upgrade Table
func (app *App) TableMigration() {
	var tables []interface{} = []interface{}{
		&model.Hotel{},
		&model.RatePlan{},
		&model.Room{},
	}

	for _, table := range tables {
		err := app.DB.AutoMigrate(table).Error
		if err != nil {
			app.Log.Errorf("Table Migration Error => %s", err.Error())
		}
	}

	if err := app.DB.Model(&model.Room{}).
		AddForeignKey("hotel_id", "hotels(hotel_id)", "RESTRICT", "RESTRICT").Error; err != nil {
		app.Log.Errorf("Foreign Key Error => %s", err.Error())
	}

	if err := app.DB.Model(&model.RatePlan{}).
		AddForeignKey("hotel_id", "hotels(hotel_id)", "RESTRICT", "RESTRICT").Error; err != nil {
		app.Log.Errorf("Foreign Key Error => %s", err.Error())
	}
}
