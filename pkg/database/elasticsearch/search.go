package elasticsearch

import (
	"context"
	"encoding/json"

	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	elastic "github.com/olivere/elastic/v7"
)

type search interface {
	doSearch() (*elastic.SearchResult, error)
	unmarshal() unmarshal
}

// unmarshal content returned after unmarshalling data
type unmarshal struct {
	content interface{}
}

type configElastic struct {
	ctx             context.Context
	client          *elastic.Client
	searchKey       string
	searchThisValue string
}

// searchEventImpl Make a search on event-store index
type searchEventImpl struct {
	*configElastic
}

func newSearchEventImpl(config *configElastic) *searchEventImpl {
	return &searchEventImpl{
		configElastic: config,
	}
}

// searchEventImpl Make a search on event-store index
func (searchEventImpl *searchEventImpl) doSearch() (*elastic.SearchResult, error) {
	elasticConfig := searchEventImpl.configElastic

	searchResult, err := elasticConfig.client.Search().
		Index(indexEventStore).
		Query(elastic.NewMatchAllQuery()).
		Do(elasticConfig.ctx)

	if err != nil {
		return &elastic.SearchResult{}, err
	}

	return searchResult, nil
}

func (searchEventImpl searchEventImpl) unmarshal(fromData *elastic.SearchResult) (unmarshal, error) {
	var events []db.Event
	for _, hit := range fromData.Hits.Hits {
		var event db.Event

		err := json.Unmarshal(hit.Source, &event)
		if err != nil {
			return unmarshal{}, err
		}

		events = append(events, event)
	}

	return unmarshal{
		content: events,
	}, nil
}

// searchArticleImpl Make a search on read-model index
type searchArticleImpl struct {
	*configElastic
}

func newSearchArticleImpl(config *configElastic) *searchArticleImpl {
	return &searchArticleImpl{
		configElastic: config,
	}
}

// searchArticleImpl Make a search on read-model index
func (searchArticleImpl *searchArticleImpl) doSearch() (*elastic.SearchResult, error) {
	elasticConfig := searchArticleImpl.configElastic

	searchResult, err := elasticConfig.client.Search().
		Index(indexReadModel).
		Query(elastic.NewMatchQuery(elasticConfig.searchKey, elasticConfig.searchThisValue)).
		Do(elasticConfig.ctx)

	if err != nil {
		return &elastic.SearchResult{}, err
	}

	return searchResult, nil
}

func (searchArticleImpl searchArticleImpl) unmarshal(fromData *elastic.SearchResult) (unmarshal, error) {
	var articles []db.Article
	for _, hit := range fromData.Hits.Hits {
		var article db.Article

		err := json.Unmarshal(hit.Source, &article)
		if err != nil {
			return unmarshal{}, err
		}

		articles = append(articles, article)
	}

	return unmarshal{
		content: articles,
	}, nil
}
