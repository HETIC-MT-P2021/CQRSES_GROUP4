package Messager

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/streadway/amqp"
	"time"
)

var (
	//RabbitMQChan is a pointer to a rabbitmq channel
	RabbitMQChan *amqp.Channel
	//RabbitMQQueue is a rabbitmq queue
	RabbitMQQueue amqp.Queue
)

//RBMQQueuecreation is a rabbitmq model
type RBMQQueuecreation struct {
	RabbitMQChan  *amqp.Channel
	RabbitMQQueue amqp.Queue
}

//RabbitMqEnv contains rabbitmq env credentials
type RabbitMqEnv struct {
	RabbitMqHost string `env:"RABBITMQ_HOST"`
	RabbitMqPort string `env:"RABBITMQ_PORT"`
	RabbitMqUser string `env:"RABBITMQ_DEFAULT_USER"`
	RabbitMqPass string `env:"RABBITMQ_DEFAULT_PASS"`
}

//ConnectToRabbitMQ is for connecting to rabbitmq
func ConnectToRabbitMQ() error {
	time.Sleep(50 * time.Second)

	cfg := RabbitMqEnv{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("failed to parse env: %v", err)
	}

	urlConn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMqPass,
		cfg.RabbitMqUser,
		cfg.RabbitMqHost,
		cfg.RabbitMqPort,
	)

	conn, err := amqp.Dial(urlConn)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	RabbitMQChan = ch
	RabbitMQQueue = q

	return nil
}
