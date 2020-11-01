package service

import (
	"github.com/lin-sel/pub-sub-rmq/app/rabbitmq"
)

// PublishService Have Publisher
type PublishService struct {
	Publisher *rabbitmq.Publisher
}

// AddData To Publisher
func (publishservice *PublishService) AddData(data []byte) error {
	return publishservice.Publisher.Publish(data)
}
