package queries_handler

import (
	"errors"
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/queries"
)

type ReadArticleQueryHandler struct{}

func (ch ReadArticleQueryHandler) Handle(command cqrs.Query) error {
	switch cmd := command.Payload().(type) {
	case *queries.ReadArticleQuery:
		fmt.Println("Handler.")
		fmt.Println(cmd)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewReadArticleQueryHandler() *ReadArticleQueryHandler {
	return &ReadArticleQueryHandler{}
}
