package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/deserialize"
)

func getTypeAndPayload(body string) (string, interface{}, error) {
	bodyMapped, err := deserialize.ToMAP(body)
	if err != nil {
		return "", "", err
	}

	return bodyMapped["event_type"].(string), bodyMapped["payload"], nil
}

// Consume Receives event and dispatch it to event handler
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
			eventType, eventPayload, err := getTypeAndPayload(string(d.Body))

			err = EventProcessor(&EventProcessorParams{
				EventType:    eventType,
				EventPayload: eventPayload,
				EventBus:     eventBus,
				StoreInDB:    true,
			})

			if err != nil {
				articleFound := !strings.Contains(err.Error(), elasticsearch.ArticleNotFoundError)
				if articleFound {
					log.Println(err)
					return
				}

				marshall, err := json.Marshal(eventPayload)
				if err != nil {
					log.Println(err)
					return
				}

				aggr, err := deserialize.ToMAP(string(marshall))
				if err != nil {
					log.Println(err)
					return
				}

				evsFromElastic, err := elasticsearch.LoadEvents(aggr["aggregate_article_id"].(string))
				if err != nil {
					log.Println(err)
					return
				}

				for _, evElastic := range evsFromElastic {
					body, err := json.Marshal(evElastic)
					if err != nil {
						log.Println(err)
						return
					}

					eventType, eventPayload, err = getTypeAndPayload(string(body))

					err = EventProcessor(&EventProcessorParams{
						EventType:    eventType,
						EventPayload: eventPayload,
						EventBus:     eventBus,
						StoreInDB:    false,
					})

					fmt.Printf("elastic event : %s\n", evElastic)

					if err != nil {
						log.Println(err)
						//return
					}

				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
