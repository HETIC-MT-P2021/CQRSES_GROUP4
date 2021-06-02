package rabbit

import (
	"github.com/streadway/amqp"
)

// Publish new event on rabbitmq
func (rabbit *RabbitRepositoryImpl) Publish(event string) error {
	err := rabbit.Chan.Publish(
		"",
		rabbit.Queue.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		})

	return err
}
