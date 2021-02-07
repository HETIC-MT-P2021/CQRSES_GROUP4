package Messager

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

type Broker struct {
	Channel    *amqp.Channel
	Connection *amqp.Connection
}

var Messaging *Broker

const (
	attemptsRabbitConnexion = 10
	waitForConnexion        = 5
)

func InitBroker() (error error) {

	// url := os.Getenv("AMQP_URL")

	for index := 1; index <= attemptsRabbitConnexion; index++ {
		conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
		if err != nil {
			log.Print(err, "Failed to connect to RabbitMQ")
			time.Sleep(waitForConnexion * time.Second)
			continue
		} else {
			log.Print("conn", conn)
			log.Print("err", err)
			Messaging.Connection = conn

			ch, err := conn.Channel()
			if err == nil {
				log.Print(err, "Failed to connect to RabbitMQ")
			}
			Messaging.Channel = ch
		}
		break
	}

	defer Messaging.Connection.Close()
	defer Messaging.Channel.Close()
	return nil
}
