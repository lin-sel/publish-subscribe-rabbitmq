package rabbitmq

import (
	"github.com/streadway/amqp"
)

// Publisher Have Channel
type Publisher struct {
	Chan *amqp.Channel
}

// NewPublisher Return New Object Of Publisher
func NewPublisher(conn *amqp.Connection) *Publisher {
	ch, _ := conn.Channel()

	_ = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	return &Publisher{
		Chan: ch,
	}
}

// Publish the Data
func (pub *Publisher) Publish(body []byte) error {
	err := pub.Chan.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	// conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()
	return err
}
