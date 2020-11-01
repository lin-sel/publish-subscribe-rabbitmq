package rabbitmq

import (
	"github.com/lin-sel/pub-sub-rmq/app/subscriber/controller"
	"github.com/streadway/amqp"
)

// Subscriber Have Chan, Queue
type Subscriber struct {
	Chan         *amqp.Channel
	Queue        amqp.Queue
	EventListner *controller.SubscriberController
}

// NewSubscriber Return New Object Of Subscriber
func NewSubscriber(conn *amqp.Connection,
	eventListner *controller.SubscriberController) (*Subscriber, error) {
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
		Chan:         ch,
		Queue:        q,
		EventListner: eventListner,
	}, nil
}

// Subscribe fetch data from queue
func (pub *Subscriber) Subscribe() {
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
			pub.EventListner.Add(d.Body)
		}
	}()
	<-forever
}
