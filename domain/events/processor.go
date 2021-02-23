package events

import (
	"strconv"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
	uuid "github.com/satori/go.uuid"
)

//------------------------------------------------------------------
// ArticleCreatedEvent
//------------------------------------------------------------------

//payloadToArticle Transform payload to article struct
//@see Action interface
func (event ArticleCreatedEvent) payloadToArticle(payload map[string]interface{}) database.Article {
	return database.Article{
		ID:          uuid.NewV4().String(),
		Title:       payload["title"].(string),
		Description: payload["description"].(string),
	}
}

//getOne article to elastic
//@see Action interface
func (event ArticleCreatedEvent) getOne() (database.Article, error) {
	return database.Article{}, nil
}

//storeReadModel An article in db
//@see Action interface
func (event ArticleCreatedEvent) storeReadModel(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: ArticleCreatedEventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
}

//storeEventToElastic in db
//@see Action interface
func (event ArticleCreatedEvent) storeEventToElastic(article database.Article) error {
	return elasticsearch.StoreReadmodel(article)
}

//------------------------------------------------------------------
// ArticleUpdatedEvent
//------------------------------------------------------------------

//payloadToArticle Transform payload to article struct
//@see Action interface
func (event ArticleUpdatedEvent) payloadToArticle(payload map[string]interface{}) database.Article {
	return database.Article{
		ID:          payload["aggregate_article_id"].(string),
		Title:       payload["title"].(string),
		Description: payload["description"].(string),
	}
}

//getOne article to elastic
//@see Action interface
func (event ArticleUpdatedEvent) getOne() (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(event.AggregateArticleID)
	if err != nil {
		return database.Article{}, err
	}

	return article, nil
}

//storeReadModel An article in db
//@see Action interface
func (event ArticleUpdatedEvent) storeReadModel(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: ArticleUpdatedEventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
}

//storeEventToElastic in db
//@see Action interface
func (event ArticleUpdatedEvent) storeEventToElastic(article database.Article) error {
	return elasticsearch.StoreReadmodel(article)
}
