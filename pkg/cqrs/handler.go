package cqrs

// CommandHandler Allows to manage Command
type CommandHandler interface {
	Handle(Command) error
}

// QueryHandler Allows to manage Query
type QueryHandler interface {
	Handle(Query) (interface{}, error)
}
