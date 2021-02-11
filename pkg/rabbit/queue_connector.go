package rabbit

import (
	"encoding/json"
)

//QueueConnector Serialize event and send to rabbitMQ queue
func QueueConnector(event interface{}) error {
	buffer, err := json.Marshal(event)
	if err != nil {
		return nil
	}

	err = ConnectToRabbitMQ()
	if err != nil {
		return err
	}

	return Rabbit.Publish(string(buffer))
}
