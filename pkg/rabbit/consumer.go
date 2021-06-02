package rabbit

import (
	"log"
	"strings"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/types"
)

func getTypeAndPayload(body string) (string, interface{}, error) {
	bodyMapped, err := types.StringToMAP(body)
	if err != nil {
		return "", "", err
	}

	return bodyMapped["event_type"].(string), bodyMapped["payload"], nil
}

// Consume Receives event and dispatch it to event handler
func (rabbit *RabbitRepositoryImpl) Consume(eventBus *event.EventBus) {
	msgs, err := rabbit.Chan.Consume(
		rabbit.Queue.Name, // queue
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

			eProcessor := &EventProcessor{
				Message:   d.Body,
				EventBus:  eventBus,
				StoreInDB: true,
			}
			consumeMSG, err := eProcessor.ApplyEventProcessor()
			if err != nil {
				articleFound := !strings.Contains(err.Error(), elasticsearch.ArticleNotFoundError)
				if articleFound {
					log.Println(err)
					return
				}

				article, err := consumeMSG.GetPayload()
				if err != nil {
					log.Println(err)
					return
				}

				aggregateArticleID := article["aggregate_article_id"].(string)
				if err = eProcessor.ApplyEvents(aggregateArticleID); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
