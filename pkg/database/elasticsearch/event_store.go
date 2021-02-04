package elasticsearch

import (
	"context"

	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

// StoreEvent stores an event for an aggregate article
func (r *ElasticRepository) StoreEvent(event db.Event) error {
	ctx := context.Background()

	_, err := r.client.Index().
		Index(indexEventStore).
		Type("article").
		Id(event.ID).
		BodyJson(event).
		Refresh("wait_for").
		Do(ctx)
	return err
}

// LoadEvents returns all events from elastic
func (r *ElasticRepository) LoadEvents() ([]db.Event, error) {
	config := &configElastic{
		ctx:    context.Background(),
		client: r.client,
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
