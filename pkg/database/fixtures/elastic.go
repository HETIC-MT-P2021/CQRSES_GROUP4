package fixtures

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

//NewEventStore Creates 3 events
func NewEventStore() error {
	event1 := database.Event{
		ID:        "d1dd9de9-7e95-4a4a-aac3-8188d8678402",
		EventType: "ArticleCreatedEvent",
		CreatedAt: "1613079361",
		Payload: database.Article{
			ID:          "c8d1981b-164e-4b00-8787-db9040c00f95",
			Title:       "Golang",
			Description: "Golang is insane",
		},
	}

	event2 := database.Event{
		ID:        "6e8556e2-67bb-44f4-86f0-e69bd1f92899",
		EventType: "ArticleUpdatedEvent",
		CreatedAt: "1613077611",
		Payload: database.Article{
			ID:          "c8d1981b-164e-4b00-8787-db9040c00f95",
			Title:       "Golang updated",
			Description: "Golang is insane updated",
		},
	}

	event3 := database.Event{
		ID:        "148845f6-21dd-4968-abc0-9e8029c8640b",
		EventType: "ArticleCreatedEvent",
		CreatedAt: "1613083302",
		Payload: database.Article{
			ID:          "ba1ffcfe-f9bf-49ff-b5de-937dc5f8e1c8",
			Title:       "Python",
			Description: "Python is insane",
		},
	}

	elasticImpl := elasticsearch.NewElasticRepository(elasticsearch.ElasticClient)
	if err := elasticImpl.StoreEvent(event1); err != nil {
		return err
	}
	if err := elasticImpl.StoreEvent(event2); err != nil {
		return err
	}
	if err := elasticImpl.StoreEvent(event3); err != nil {
		return err
	}

	return nil
}

//NewReadModel Creates 2 read-models
func NewReadModel() error {
	article1 := database.Article{
		ID:          "c8d1981b-164e-4b00-8787-db9040c00f95",
		Title:       "Golang updated",
		Description: "Golang is insane updated",
	}

	article2 := database.Article{
		ID:          "ba1ffcfe-f9bf-49ff-b5de-937dc5f8e1c8",
		Title:       "Python",
		Description: "Python is insane",
	}

	elasticImpl := elasticsearch.NewElasticRepository(elasticsearch.ElasticClient)
	if err := elasticImpl.StoreReadmodel(article1); err != nil {
		return err
	}
	if err := elasticImpl.StoreReadmodel(article2); err != nil {
		return err
	}
	return nil
}
