package cqrs

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
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
	typeName := pkg.TypeOf(Query)
	if _, ok := bus.handlers[typeName]; ok {
		return fmt.Errorf("duplicate Query handler registration with Query bus for Query of type: %s", typeName)
	}

	bus.handlers[typeName] = handler

	return nil
}

// Dispatch Calls good Query process
func (bus *QueryBus) Dispatch(Query Query) (interface{}, error) {
	if handler, ok := bus.handlers[Query.Type()]; ok {
		return handler.Handle(Query)
	}
	//return []database.Article{}, fmt.Errorf("the Query bus does not have a handler for Querys of type: %s", Query.Type())
	return []database.Article{}, fmt.Errorf("the Query bus does not have a handler for Querys of type: %s", Query.Type())
}

// GetLength of registred command
func (bus *QueryBus) GetLength() int {
	return len(bus.handlers)
}

// GetQueriesName of registred command
func (bus *QueryBus) GetQueriesName() []string {
	cmdsName := []string{}
	for commandName := range bus.handlers {
		cmdsName = append(cmdsName, commandName)
	}

	return cmdsName
}


// QueryImpl Overrides Query
type QueryImpl struct {
	Query interface{}
}

// NewQueryImpl Initialize an Query implementation
func NewQueryImpl(Query interface{}) *QueryImpl {
	return &QueryImpl{
		Query: Query,
	}
}

// Type Returns query type
func (c *QueryImpl) Type() string {
	return pkg.TypeOf(c.Query)
}

// Payload returns the actual Query payload of the message.
func (c *QueryImpl) Payload() interface{} {
	return c.Query
}
