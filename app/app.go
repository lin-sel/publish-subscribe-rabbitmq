package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	pubcon "github.com/lin-sel/pub-sub-rmq/app/publisher/controller"
	pubser "github.com/lin-sel/pub-sub-rmq/app/publisher/service"
	"github.com/lin-sel/pub-sub-rmq/app/rabbitmq"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/controller"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/model"
	"github.com/sirupsen/logrus"
)

// App DB, Log
type App struct {
	DB    *gorm.DB
	Log   *logrus.Logger
	Route *mux.Router
}

// NewApp Return New Object Of App
func NewApp(db *gorm.DB, lg *logrus.Logger, route *mux.Router) *App {
	return &App{
		DB:    db,
		Log:   lg,
		Route: route,
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

// RabbitMQConfig Config Messanger
func (app *App) RabbitMQConfig(contr *controller.SubscriberController) *rabbitmq.Publisher {
	conn, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		app.Log.Error(err.Error())
	}
	publisher := rabbitmq.NewPublisher(conn)
	subcriber, err := rabbitmq.NewSubscriber(conn, contr)
	if err != nil {
		app.Log.Error(err.Error())
	}
	go subcriber.Subscribe()
	return publisher
}

// InitApp App
func (app *App) InitApp(contr *controller.SubscriberController) {
	pub := app.RabbitMQConfig(contr)
	app.RegisterPublisher(pub)
	app.TableMigration()
	if err := http.ListenAndServe(":8080", app.Route); err != nil {
		app.Log.Fatal(err.Error())
	}
}

// RegisterPublisher Register Publisher
func (app *App) RegisterPublisher(pub *rabbitmq.Publisher) {
	pubservice := pubser.NewPublishService(pub)
	pubcontroller := pubcon.NewPublishController(pubservice)
	pubcontroller.RegisterRoute(app.Route)
}
