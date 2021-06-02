package rabbit

import (
	"encoding/json"
)

// QueueConnector Serialize event and send to rabbitMQ queue
func (rabbit *RabbitRepositoryImpl) QueueConnector(event interface{}) error {
	buffer, err := json.Marshal(event)
	if err != nil {
		return nil
	}

	return rabbit.Publish(string(buffer))
}
