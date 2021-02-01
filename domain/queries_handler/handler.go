package queries_handler

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

type ReadArticlesQueryHandler struct{}

func (ch ReadArticlesQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	switch qu := query.Payload().(type) {
	case *queries.ReadArticlesQuery:
		fmt.Println(qu)
		fmt.Println("fdp ntm")
		//fmt.Println(state.CurrentArticles.Articles())
		return []database.Article{}, nil
	default:
		return []database.Article{}, errors.New("bad command type")
	}
}

func NewReadArticlesQueryHandler() *ReadArticlesQueryHandler {
	return &ReadArticlesQueryHandler{}
}
