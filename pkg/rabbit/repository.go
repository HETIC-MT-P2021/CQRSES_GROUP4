package rabbit

import "github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"

type RabbitRepository interface {
	QueueConnector(interface{}) error
	Publish(string) error
	Consume(*event.EventBus)
}