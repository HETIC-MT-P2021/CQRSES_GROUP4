package queries_handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

type ReadArticlesQueryHandler struct{}

func (ch ReadArticlesQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	switch qu := query.Payload().(type) {
	case *queries.ReadArticlesQuery:
		fmt.Println(qu)

		// pass aggregate_article_id
		articles, err := elasticsearch.GetReadmodel("")
		if err != nil {
			log.Println(err)
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
