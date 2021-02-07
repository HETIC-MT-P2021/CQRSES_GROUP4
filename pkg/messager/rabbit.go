package messager

import "github.com/streadway/amqp"

//RabbitMQ Allows to make process on a rabbitConnector type
type RabbitMQ interface {
	Publish() error
	Consume()
}

//rabbitConnector stores channel and queue to make publish possible
type rabbitConnector struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

//newRabbitConnector Creates instance of rabbitConnector
func newRabbitConnector(channel *amqp.Channel, queue amqp.Queue) *rabbitConnector {
	return &rabbitConnector{
		Chan:  channel,
		Queue: queue,
	}
}
