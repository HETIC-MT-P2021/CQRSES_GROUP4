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

//update article state
//@see Action interface
func (event ArticleCreatedEvent) update(articlePayload map[string]interface{}) (database.Article, error) {
	return database.Article{
		ID:          articlePayload["aggregate_article_id"].(string),
		Title:       articlePayload["title"].(string),
		Description: articlePayload["description"].(string),
	}, nil
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

//update article state
//@see Action interface
func (event ArticleUpdatedEvent) update(articlePayload map[string]interface{}) (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(event.AggregateArticleID)
	if err != nil {
		return database.Article{}, err
	}

	article.Title = articlePayload["title"].(string)
	article.Description = articlePayload["description"].(string)

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
