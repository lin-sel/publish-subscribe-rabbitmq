package service

import (
	"net/http"

	"github.com/lin-sel/pub-sub-rmq/app/util"

	"github.com/jinzhu/gorm"
	errors "github.com/lin-sel/pub-sub-rmq/app/error"
	"github.com/lin-sel/pub-sub-rmq/app/repository"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/model"
)

// HotelService Contain DB, Repository
type HotelService struct {
	DB         *gorm.DB
	Repository repository.Repository
}

// NewHotelService Return New Object Of HotelService
func NewHotelService(db *gorm.DB, repo repository.Repository) *HotelService {
	return &HotelService{
		DB:         db,
		Repository: repo,
	}
}

// AddHotel Add New Hotel Record
func (roomservice *HotelService) AddHotel(hotel *model.Hotel) error {
	uow := repository.NewUnitOfWork(roomservice.DB, false)
	hotel.ID = util.GenerateUUID()
	err := roomservice.Repository.Add(uow, hotel)
	if err != nil {
		uow.Complete()
		return errors.NewHTTPError(err.Error(), http.StatusInternalServerError)
	}
	uow.Commit()
	return nil
}
