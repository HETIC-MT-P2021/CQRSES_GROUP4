package main

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
)

func main() {
	domain.InitBusses()

	err := rabbit.ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}
	rabbit.Rabbit.Consume(domain.EventBus)
}
