package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

type Payload struct {
	Content interface{}
}

func deserializeToMAP(serialized string) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(serialized), &data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}

// Consume test
func (connector *RabbitRepository) Consume(eventBus *event.EventBus) {
	msgs, err := connector.Chan.Consume(
		connector.Queue.Name, // queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Println(err)
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			dataMapped, err := deserializeToMAP(string(d.Body))
			if err != nil {
				log.Println(err)
			}

			fmt.Println(dataMapped["event_type"])

			switch dataMapped["event_type"] {
			case "ArticleCreatedEvent":
				buffer, err := json.Marshal(dataMapped["payload"])
				if err != nil {
					return
				}

				payload := string(buffer)
				payloadMapped, err := deserializeToMAP(payload)
				if err != nil {
					log.Println(err)
				}

				ev := &events.ArticleCreatedEvent{
					ID:          payloadMapped["event_id"].(string),
					Title:       payloadMapped["title"].(string),
					Description: payloadMapped["description"].(string),
				}

				eventImpl := event.NewEventImpl(ev)
				err = eventBus.Dispatch(eventImpl)
				if err != nil {
					log.Println(err)
				}

			case "ArticleUpdatedEvent":
			default:
				return
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
