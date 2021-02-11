package rabbit

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

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
			err := EventProcessor(&EventProcessorParams{
				Body:     d.Body,
				EventBus: eventBus,
			})

			if err != nil {
				log.Println(err)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
