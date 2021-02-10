package rabbit

import (
	"github.com/streadway/amqp"
)

//Publish new event on rabbitmq
func (connector *RabbitRepository) Publish(event string) error {
	err := connector.Chan.Publish(
		"",
		connector.Queue.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		})

	return err
}
