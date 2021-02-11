package article

import "github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"

//ActionRequested To do on article
type ActionRequested interface {
	GetOne() (database.Article, error)
	Store(database.Article) error
}

//Create new Article
type Create struct {
	EventType string
}

//Update new Article
type Update struct {
	AggregateArticleID string
	EventType          string
}
