package service

import (
	"net/http"

	"github.com/jinzhu/gorm"
	errors "github.com/lin-sel/pub-sub-rmq/error"
	"github.com/lin-sel/pub-sub-rmq/model"
	"github.com/lin-sel/pub-sub-rmq/repository"
)

// HotelService Contain DB, Repository
type HotelService struct {
	DB         *gorm.DB
	Repository repository.Repository
}

// AddHotel Add New Hotel Record
func (roomservice *HotelService) AddHotel(hotel *model.Hotel) error {
	uow := repository.NewUnitOfWork(roomservice.DB, false)

	err := roomservice.Repository.Add(uow, hotel)
	if err != nil {
		return errors.NewHTTPError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
