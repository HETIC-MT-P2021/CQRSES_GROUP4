package rabbit

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/deserialize"
)

//ConsumeMessage Send message to rabbitMQ
type ConsumeMessage struct {
	EventType string      `json:"event_type"`
	Payload   interface{} `json:"payload"`
}

//GetPayload Map payload as map[string]interface{} type
func (message ConsumeMessage) GetPayload() (map[string]interface{}, error) {
	marshall, err := json.Marshal(message.Payload)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return deserialize.ToMAP(string(marshall))
}
