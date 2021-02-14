package article

import (
	"strconv"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	uuid "github.com/satori/go.uuid"
)

//PayloadToArticle Transform payload to article struct
//@see ActionRequested interface
func (create Create) PayloadToArticle(payload map[string]interface{}) database.Article {
	return database.Article{
		ID:          payload["aggregate_article_id"].(string),
		Title:       payload["title"].(string),
		Description: payload["description"].(string),
	}
}

//GetOne article to elastic
//@see ActionRequested interface
func (create Create) GetOne() (database.Article, error) {
	return database.Article{}, nil
}

//StoreEventToElastic in db
//@see ActionRequested interface
func (create Create) StoreEventToElastic(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: create.EventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
}

//Store An article to create
//@see ActionRequested interface
func (create Create) Store(article database.Article) error {
	return elasticsearch.StoreReadmodel(article)
}
