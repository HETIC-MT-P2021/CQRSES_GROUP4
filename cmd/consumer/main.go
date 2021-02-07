package main

import (
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/messager"
)

func main() {
	connector, err := messager.ConnectToRabbitMQ()
	if err != nil {
		log.Println(err)
		return
	}

	connector.Consume()
}
