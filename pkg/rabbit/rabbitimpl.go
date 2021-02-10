package rabbit

import (
	"fmt"
	"time"

	env "github.com/caarlos0/env"
	"github.com/streadway/amqp"
)

type Repository interface {
	Publish(string) error
	Consume()
}

type RabbitRepository struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

var Rabbit *RabbitRepository

//rabbitMqEnv contains rabbitmq env credentials
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

//ConnectToRabbitMQ is for connecting to rabbitmq
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

	Rabbit = &RabbitRepository{
		Chan:  ch,
		Queue: q,
	}

	return nil
}
