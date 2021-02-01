package state

import "github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"

var CurrentArticles *event.ArticlesAggregate

func InitState() {
	CurrentArticles = event.NewEmpty()
}
