package elasticsearch

import (
	"context"
	"encoding/json"

	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	elastic "github.com/olivere/elastic/v7"
)

type search struct {
	Ctx             context.Context
	Client          *elastic.Client
	SearchKey       string
	SearchThisValue string
}

func (search search) initSearch() (*elastic.SearchSource, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("event_name", "ArticleCreatedEvent"))

	_, err := searchSource.Source()
	if err != nil {
		return &elastic.SearchSource{}, err
	}

	return searchSource, nil
}

func (search search) doSearch(searchSource *elastic.SearchSource) (*elastic.SearchResult, error) {
	searchResult, err := search.Client.Search().Index("hetic").
		SearchSource(searchSource).
		Do(search.Ctx)

	if err != nil {
		return &elastic.SearchResult{}, err
	}

	return searchResult, nil
}

func (search search) unmarshallEvents(fromData *elastic.SearchResult) ([]db.Event, error) {
	var events []db.Event
	for _, hit := range fromData.Hits.Hits {
		var event db.Event

		err := json.Unmarshal(hit.Source, &event)
		if err != nil {
			return []db.Event{}, err
		}

		events = append(events, event)
	}

	return events, nil
}
