package messager

import (
	"github.com/streadway/amqp"
)

//Publish new event on rabbitmq
func (connector rabbitConnector) Publish() error {
	body := "1"
	err := connector.Chan.Publish(
		"",
		connector.Queue.Name, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	return err
}
