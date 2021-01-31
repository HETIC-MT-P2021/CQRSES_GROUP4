package elasticsearch

import (
	"context"
	"fmt"

	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
	elastic "github.com/olivere/elastic/v7"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	fmt.Println("url printed " + url)

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	if err != nil {
		return nil, err
	}

	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) StoreEvent(event db.Event) error {
	ctx := context.Background()

	_, err := r.client.Index().
		Index("hetic").
		Type("events").
		Id(event.ID).
		BodyJson(event).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) LoadEvents() ([]db.Event, error) {
	search := search{
		Ctx:             context.Background(),
		Client:          r.client,
		SearchKey:       "event_name",
		SearchThisValue: "ArticleCreatedEvent",
	}
	searchSource, err := search.initSearch()
	if err != nil {
		return []db.Event{}, err
	}

	searchResult, err := search.doSearch(searchSource)
	if err != nil {
		return []db.Event{}, err
	}

	events, err := search.unmarshallEvents(searchResult)
	if err != nil {
		return []db.Event{}, err
	}

	return events, nil
}
