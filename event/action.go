package event

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

//EventApplyer To make to event
type EventApplyer interface {
	Apply(Event) error

	payloadToArticle(map[string]interface{}) database.Article
	getOne(map[string]interface{}) (database.Article, error)
	storeReadModel(database.Article) error
	storeEventToElastic(database.Article) error
}
