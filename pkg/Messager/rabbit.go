package Messager

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type Broker struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

var Messaging *Broker

const (
	attemptsRabbitConnexion = 3
	waitForConnexion        = 3
)

func InitBroker() (error error) {

	url := os.Getenv("AMQP_URL")

	for index := 1; index <= attemptsRabbitConnexion; index++ {
		connection, err := amqp.Dial(url)
		if err != nil {
			if index < attemptsRabbitConnexion {
				log.Printf("Rabbit connection, %d retry : %v", index, err)
				time.Sleep(waitForConnexion * time.Second)
			}
			continue
		} else {
			Messaging.Connection = connection

			channel, err := connection.Channel()

			if err != nil {
				log.Print(err)
				return err
			}

			Messaging.Channel = channel
		}

		break
	}
	defer Messaging.Connection.Close()

	return nil
}
