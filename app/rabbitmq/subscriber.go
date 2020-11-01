package rabbitmq

import (
	"github.com/streadway/amqp"
)

// Subscriber Have Chan, Queue
type Subscriber struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

// NewSubscriber Return New Object Of Subscriber
func (pub *Subscriber) NewSubscriber(conn *amqp.Connection) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	err = ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		"logs", // exchange
		false,
		nil,
	)
	return &Subscriber{
		Chan:  ch,
		Queue: q,
	}, nil
}

// GetMessages Return Message
func (pub *Subscriber) GetMessages(ch chan<- []byte) {
	msgs, _ := pub.Chan.Consume(
		pub.Queue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			ch <- d.Body
		}
	}()
	<-forever
}
