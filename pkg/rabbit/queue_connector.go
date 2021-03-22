package rabbit

import (
	"encoding/json"
	"fmt"
)

//QueueConnector Serialize event and send to rabbitMQ queue
func QueueConnector(event interface{}) error {
	buffer, err := json.Marshal(event)
	if err != nil {
		return nil
	}

	fmt.Println("event data2")
	fmt.Println(string(buffer))

	err = ConnectToRabbitMQ()
	if err != nil {
		return err
	}

	return Rabbit.Publish(string(buffer))
}
