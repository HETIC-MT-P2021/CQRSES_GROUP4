package queries

import (
	"errors"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

// ReadArticleQueryHandler allows to get article
type ReadArticleQueryHandler struct{
	Repo elasticsearch.ElasticRepository
}

// Handle Get an article from elasticsearch database
func (qHandler ReadArticleQueryHandler) Handle(query cqrs.Query) (interface{}, error) {
	switch query.Payload().(type) {
	case *ReadArticleQuery:
		aggregateArticleID := query.Payload().(*ReadArticleQuery).AggregateArticleID

		if aggregateArticleID == "" {
			return []database.Article{}, errors.New("aggregateArticleID should not be empty")
		}

		articles, err := qHandler.Repo.GetReadmodel(aggregateArticleID)
		if err != nil {
			return []database.Article{}, err
		}

		return articles, nil
	default:
		return []database.Article{}, errors.New("bad command type")
	}
}

// NewReadArticleQueryHandler Creates an instance
func NewReadArticleQueryHandler(repo elasticsearch.ElasticRepository) *ReadArticleQueryHandler {
	return &ReadArticleQueryHandler{
		Repo: repo,
	}
}
