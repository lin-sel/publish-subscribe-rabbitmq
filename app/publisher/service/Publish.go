package service

import (
	"github.com/lin-sel/pub-sub-rmq/app/rabbitmq"
)

// PublishService Have Publisher
type PublishService struct {
	Publisher *rabbitmq.Publisher
}

// NewPublishService Return New Object Of PublishController
func NewPublishService(pub *rabbitmq.Publisher) *PublishService {
	return &PublishService{
		Publisher: pub,
	}
}

// AddData To Publisher
func (publishservice *PublishService) AddData(data []byte) error {
	return publishservice.Publisher.Publish(data)
}
