package service

import (
	"net/http"

	"github.com/lin-sel/pub-sub-rmq/app/util"

	"github.com/jinzhu/gorm"
	errors "github.com/lin-sel/pub-sub-rmq/app/error"
	"github.com/lin-sel/pub-sub-rmq/app/repository"
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/model"
)

// RatePlanService Contain DB, Repository
type RatePlanService struct {
	DB         *gorm.DB
	Repository repository.Repository
}

// NewRatePlanService Return New Object Of RatePlanService
func NewRatePlanService(db *gorm.DB, repo repository.Repository) *RatePlanService {
	return &RatePlanService{
		DB:         db,
		Repository: repo,
	}
}

// AddRatePlan Add New RatePlan Record
func (rateplanservice *RatePlanService) AddRatePlan(room *model.RatePlan) error {
	uow := repository.NewUnitOfWork(rateplanservice.DB, false)
	room.ID = util.GenerateUUID()
	err := rateplanservice.Repository.Add(uow, room)
	if err != nil {
		uow.Complete()
		return errors.NewHTTPError(err.Error(), http.StatusInternalServerError)
	}
	uow.Commit()
	return nil
}
