package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	errors "github.com/lin-sel/pub-sub-rmq/app/error"
	"github.com/lin-sel/pub-sub-rmq/app/publisher/service"
	"github.com/lin-sel/pub-sub-rmq/app/web"
)

// PublishController Have PublishService
type PublishController struct {
	PublishService service.PublishService
}

// NewPublishController Return New Object Of PublishController
func NewPublishController(ser service.PublishService) *PublishController {
	return &PublishController{
		PublishService: ser,
	}
}

// RegisterRoute Register Endpoints
func (pub *PublishController) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/publish", pub.AddData).Methods(http.MethodPost)
}

// AddData Add Post To Data
func (pub *PublishController) AddData(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		web.RespondError(w, errors.NewHTTPError("request body has empty", http.StatusNoContent))
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		web.RespondError(w, errors.NewHTTPError("unable to read request body", http.StatusInternalServerError))
	}
	err = pub.PublishService.AddData(requestBody)
	if err != nil {
		web.RespondError(w, errors.NewHTTPError(err.Error(), http.StatusInternalServerError))
	}
}
