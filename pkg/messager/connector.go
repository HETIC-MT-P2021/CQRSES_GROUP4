package messager

import (
	"fmt"
	"time"

	env "github.com/caarlos0/env"
	"github.com/streadway/amqp"
)

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
func ConnectToRabbitMQ() (*rabbitConnector, error) {
	cfg := rabbitMqEnv{}
	if err := env.Parse(&cfg); err != nil {
		return &rabbitConnector{}, fmt.Errorf("failed to parse env: %v", err)
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
		return &rabbitConnector{}, fmt.Errorf("failed to open a channel: %v", err)
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
		return &rabbitConnector{}, fmt.Errorf("failed to declare a queue: %v", err)
	}

	return newRabbitConnector(ch, q), nil
}
