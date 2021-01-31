package cqrs

import (
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/helper"
)

// CommandBus Contains handlers
type CommandBus struct {
	handlers map[string]CommandHandler
}

// Command Create custom Command
type Command interface {
	Payload() interface{}
	Type() string
}

// NewCommandBus Initialize empty handlers in bus
func NewCommandBus() *CommandBus {
	cBus := &CommandBus{
		handlers: make(map[string]CommandHandler),
	}

	return cBus
}

// AddHandler to bus
func (bus *CommandBus) AddHandler(handler CommandHandler, command interface{}) error {
	typeName := helper.TypeOf(command)
	if _, ok := bus.handlers[typeName]; ok {
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	bus.handlers[typeName] = handler

	return nil
}

// Dispatch Calls good command process
func (bus *CommandBus) Dispatch(command Command) error {
	if handler, ok := bus.handlers[command.Type()]; ok {
		return handler.Handle(command)
	}
	return fmt.Errorf("the command bus does not have a handler for commands of type: %s", command.Type())
}

// CommandImpl Overrides Command
type CommandImpl struct {
	command interface{}
}

// NewCommand Initialize an Command implementation
func NewCommand(command interface{}) *CommandImpl {
	return &CommandImpl{
		command: command,
	}
}

// Type Returns event type
func (c *CommandImpl) Type() string {
	return helper.TypeOf(c.command)
}

// Payload returns the actual command payload of the message.
func (c *CommandImpl) Payload() interface{} {
	return c.command
}
