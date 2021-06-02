package elasticsearch

import (
	"context"

	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

// StoreEvent stores an event for an aggregate article
func (r *ElasticRepositoryImpl) StoreEvent(event db.Event) error {
	ctx := context.Background()

	_, err := ElasticClient.Index().
		Index(indexEventStore).
		Id(event.ID).
		BodyJson(event).
		Refresh("wait_for").
		Do(ctx)

	return err
}

// LoadEvents returns all events from elastic
func (r *ElasticRepositoryImpl) LoadEvents(aggregateArticleID string) ([]db.Event, error) {
	config := &configElastic{
		ctx:             context.Background(),
		client:          ElasticClient,
		searchKey:       "payload.aggregate_article_id",
		searchThisValue: aggregateArticleID,
	}

	searchEventImpl := newSearchEventImpl(config)
	searchResult, err := searchEventImpl.doSearch()
	if err != nil {
		return []db.Event{}, err
	}

	events, err := searchEventImpl.unmarshal(searchResult)
	if err != nil {
		return []db.Event{}, err
	}

	return events.content.([]db.Event), nil
}
