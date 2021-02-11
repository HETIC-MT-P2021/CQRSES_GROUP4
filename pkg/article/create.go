package article

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/elasticsearch"
)

//GetOne article to elastic
//@see ActionRequested interface
func (create Create) GetOne() (database.Article, error) {
	return database.Article{}, nil
}

//Store An article to create
//@see ActionRequested interface
func (create Create) Store(article database.Article) error {
	if err := storeEventToElastic(create.EventType, article); err != nil {
		return err
	}

	return elasticsearch.StoreReadmodel(article)
}
