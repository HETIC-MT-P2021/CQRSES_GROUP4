package queries_handler

import (
	"errors"
	"fmt"

	"github.com/jibe0123/CQRSES_GROUP4/pkg/cqrs"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/queries"
	"github.com/jibe0123/CQRSES_GROUP4/pkg/domain/state"
)

type ReadArticlesQueryHandler struct{}

func (ch ReadArticlesQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	switch qu := query.Payload().(type) {
	case *queries.ReadArticlesQuery:
		fmt.Println(qu)
		fmt.Println("fdp ntm")
		fmt.Println(state.CurrentArticles.Articles())
		return []database.Article{}, nil
	default:
		return []database.Article{}, errors.New("bad command type")
	}
}

func NewReadArticlesQueryHandler() *ReadArticlesQueryHandler {
	return &ReadArticlesQueryHandler{}
}
