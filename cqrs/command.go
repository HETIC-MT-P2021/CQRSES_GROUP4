package cqrs

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
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
	typeName := pkg.TypeOf(command)
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

// NewCommandImpl Initialize an Command implementation
func NewCommandImpl(command interface{}) *CommandImpl {
	return &CommandImpl{
		command: command,
	}
}

// Type Returns event type
func (c *CommandImpl) Type() string {
	return pkg.TypeOf(c.command)
}

// Payload returns the actual command payload of the message.
func (c *CommandImpl) Payload() interface{} {
	return c.command
}
