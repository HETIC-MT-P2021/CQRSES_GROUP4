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
func (update Update) PayloadToArticle(payload map[string]interface{}) database.Article {
	return database.Article{
		ID:          payload["aggregate_article_id"].(string),
		Title:       payload["title"].(string),
		Description: payload["description"].(string),
	}
}

//GetOne article to elastic
//@see ActionRequested interface
func (update Update) GetOne() (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(update.AggregateArticleID)
	if err != nil {
		return database.Article{}, err
	}

	return article, nil
}

//StoreEventToElastic in db
//@see ActionRequested interface
func (update Update) StoreEventToElastic(article database.Article) error {
	createdAt := strconv.FormatInt(time.Now().Unix(), 10)
	newEvent := database.Event{
		ID:        uuid.NewV4().String(),
		EventType: update.EventType,
		CreatedAt: createdAt,
		Payload:   article,
	}

	return elasticsearch.StoreEvent(newEvent)
}

//Store An article to update to elastic
//@see ActionRequested interface
func (update Update) Store(article database.Article) error {
	return elasticsearch.UpdateReadmodel(update.AggregateArticleID, article)
}
