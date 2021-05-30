package elasticsearch

import (
	"context"
	"errors"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	db "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

// StoreReadmodel stores a readmodel for an aggregate article
func (r *ElasticRepository) StoreReadmodel(article database.Article) error {
	ctx := context.Background()

	_, err := ElasticClient.Index().
		Index(indexReadModel).
		Type("article").
		Id(article.ID).
		BodyJson(article).
		Refresh("wait_for").
		Do(ctx)
	return err
}

// UpdateReadmodel stores an updated readmodel for an aggregate article
func (r *ElasticRepository) UpdateReadmodel(aggregateArticleID string,
	article database.Article) error {
	ctx := context.Background()

	_, err := ElasticClient.Update().
		Index(indexReadModel).
		Type("article").
		Id(aggregateArticleID).
		Doc(article).
		Do(ctx)
	return err
}

// GetReadmodel returns an article from elastic
// pass aggregateID as param to get an article
func (r *ElasticRepository) GetReadmodel(aggregateID string) (db.Article, error) {
	config := &configElastic{
		ctx:             context.Background(),
		client:          ElasticClient,
		searchKey:       "_id",
		searchThisValue: aggregateID,
	}

	searchArticleImpl := newSearchArticleImpl(config)
	searchResult, err := searchArticleImpl.doSearch()
	if err != nil {
		return db.Article{}, err
	}

	article, err := searchArticleImpl.unmarshal(searchResult)
	if err != nil {
		return db.Article{}, err
	}

	articleIsEmpty := len(article.content.([]db.Article)) <= 0
	if articleIsEmpty {
		return db.Article{}, errors.New(ArticleNotFoundError)
	}

	return article.content.([]db.Article)[0], nil
}
