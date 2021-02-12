package article

import (
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

//GetOne article to elastic
//@see ActionRequested interface
func (update Update) GetOne() (database.Article, error) {
	article, err := elasticsearch.GetReadmodel(update.AggregateArticleID)
	if err != nil {
		fmt.Println("here is error")
		log.Println(err)
		return database.Article{}, err
	}

	fmt.Println("article found")
	fmt.Println(article)

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
