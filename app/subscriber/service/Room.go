package service

import (
	"net/http"

	"github.com/jinzhu/gorm"
	errors "github.com/lin-sel/pub-sub-rmq/app/error"
	"github.com/lin-sel/pub-sub-rmq/app/repository"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/model"
)

// RoomService Contain DB, Repository
type RoomService struct {
	DB         *gorm.DB
	Repository repository.Repository
}

// NewRoomService Return New Object Of RoomService
func NewRoomService(db *gorm.DB, repo repository.Repository) *RoomService {
	return &RoomService{
		DB:         db,
		Repository: repo,
	}
}

// AddRoom Add New Room Record
func (roomservice *RoomService) AddRoom(room *model.Room) error {
	uow := repository.NewUnitOfWork(roomservice.DB, false)

	err := roomservice.Repository.Add(uow, room)
	if err != nil {
		return errors.NewHTTPError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
