package types

import (
	"encoding/json"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

func GetPayloadMapped(ev event.Event) (map[string]interface{}, error) {
	eventPayload := ev.Payload()
	buffer, err := json.Marshal(eventPayload)
	if err != nil {
		return map[string]interface{}{}, err
	}

	payload := string(buffer)
	return StringToMAP(payload)
}