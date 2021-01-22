package main

import "github.com/jibe0123/CQRSES_GROUP4/pkg/rabbitmq"

// Start consumer for waiting event
func main() {
	rabbitmq.ReceiveEventMessage()
}
