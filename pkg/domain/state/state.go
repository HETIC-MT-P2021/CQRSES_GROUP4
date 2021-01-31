package state

import "github.com/jibe0123/CQRSES_GROUP4/pkg/event"

var CurrentArticles *event.ArticlesAggregate

func InitState() {
	CurrentArticles = event.NewEmpty()
}
