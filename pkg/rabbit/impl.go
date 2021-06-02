package rabbit

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/streadway/amqp"
)

var RabbitChannel *amqp.Channel
var RabbitQueue amqp.Queue

type RabbitRepositoryImpl struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

func NewRabbitRepository(channel *amqp.Channel, queue amqp.Queue) *RabbitRepositoryImpl {
	return &RabbitRepositoryImpl{
		Chan: channel,
		Queue: queue,
	}
}

// rabbitMqEnv contains rabbitmq env credentials
type rabbitMqEnv struct {
	RabbitMqHost string `env:"RABBITMQ_HOST"`
	RabbitMqPort string `env:"RABBITMQ_PORT"`
	RabbitMqUser string `env:"RABBITMQ_DEFAULT_USER"`
	RabbitMqPass string `env:"RABBITMQ_DEFAULT_PASS"`
}

const (
	numberOftries       = 10
	timeToWaitInSeconds = 5
)

// ConnectToRabbitMQ is for connecting to rabbitmq
func ConnectToRabbitMQ() error {
	cfg := rabbitMqEnv{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	urlConn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMqPass,
		cfg.RabbitMqUser,
		cfg.RabbitMqHost,
		cfg.RabbitMqPort)

	var rabbitConnection *amqp.Connection
	var err error
	for index := 0; index < numberOftries; index++ {
		rabbitConnection, err = amqp.Dial(urlConn)
		if err != nil {
			time.Sleep(timeToWaitInSeconds * time.Second)
		} else {
			break
		}
	}

	ch, err := rabbitConnection.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"event",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	RabbitChannel = ch
	RabbitQueue = q

	return nil
}
