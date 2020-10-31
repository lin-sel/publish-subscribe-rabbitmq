package service

import (
	"net/http"

	"github.com/jinzhu/gorm"
	errors "github.com/lin-sel/pub-sub-rmq/error"
	"github.com/lin-sel/pub-sub-rmq/model"
	"github.com/lin-sel/pub-sub-rmq/repository"
)

// RatePlanService Contain DB, Repository
type RatePlanService struct {
	DB         *gorm.DB
	Repository repository.Repository
}

// AddRatePlan Add New RatePlan Record
func (rateplanservice *RatePlanService) AddRatePlan(room *model.RatePlan) error {
	uow := repository.NewUnitOfWork(rateplanservice.DB, false)

	err := rateplanservice.Repository.Add(uow, room)
	if err != nil {
		return errors.NewHTTPError(err.Error(), http.StatusInternalServerError)
	}
	return nil
}
