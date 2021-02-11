package article

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

//GetOne article to elastic
//@see ActionRequested interface
func (update Update) GetOne() (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(update.AggregateArticleID)
	if err != nil {
		return database.Article{}, err
	}

	return article, nil
}

//Store An article to update to elastic
//@see ActionRequested interface
func (update Update) Store(article database.Article) error {
	if err := storeEventToElastic(update.EventType, article); err != nil {
		return err
	}

	return elasticsearch.UpdateReadmodel(update.AggregateArticleID, article)
}
