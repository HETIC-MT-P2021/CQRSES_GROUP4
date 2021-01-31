package cqrs

import (
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/helper"
)

// QueryBus Contains handlers
type QueryBus struct {
	handlers map[string]QueryHandler
}

// Query Create custom Query
type Query interface {
	Payload() interface{}
	Type() string
}

// NewQueryBus Initialize empty handlers in bus
func NewQueryBus() *QueryBus {
	cBus := &QueryBus{
		handlers: make(map[string]QueryHandler),
	}

	return cBus
}

// AddHandler to bus
func (bus *QueryBus) AddHandler(handler QueryHandler, Query interface{}) error {
	typeName := helper.TypeOf(Query)
	if _, ok := bus.handlers[typeName]; ok {
		return fmt.Errorf("duplicate Query handler registration with Query bus for Query of type: %s", typeName)
	}

	bus.handlers[typeName] = handler

	return nil
}

// Dispatch Calls good Query process
func (bus *QueryBus) Dispatch(Query Query) error {
	if handler, ok := bus.handlers[Query.Type()]; ok {
		return handler.Handle(Query)
	}
	return fmt.Errorf("the Query bus does not have a handler for Querys of type: %s", Query.Type())
}

// QueryImpl Overrides Query
type QueryImpl struct {
	Query interface{}
}

// NewQuery Initialize an Query implementation
func NewQuery(Query interface{}) *QueryImpl {
	return &QueryImpl{
		Query: Query,
	}
}

// Type Returns event type
func (c *QueryImpl) Type() string {
	return helper.TypeOf(c.Query)
}

// Payload returns the actual Query payload of the message.
func (c *QueryImpl) Payload() interface{} {
	return c.Query
}
