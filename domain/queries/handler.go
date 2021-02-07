package queries

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

type ReadArticlesQueryHandler struct{}

func (qHandler ReadArticlesQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	switch qu := query.Payload().(type) {
	case *ReadArticleQuery:
		fmt.Println(qu)

		aggregateArticleID := query.Payload().(*ReadArticleQuery).AggregateArticleID
		articles, err := elasticsearch.GetReadmodel(aggregateArticleID)
		if err != nil {
			return []database.Article{}, err
		}

		return articles, nil
	default:
		return []database.Article{}, errors.New("bad command type")
	}
}

func NewReadArticlesQueryHandler() *ReadArticlesQueryHandler {
	return &ReadArticlesQueryHandler{}
}
