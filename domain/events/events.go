package events

//ArticleCreatedEventType is an event
var ArticleCreatedEventType = "ArticleCreatedEvent"

//ArticleUpdatedEventType is an event
var ArticleUpdatedEventType = "ArticleUpdatedEvent"

//ArticleCreatedEvent Event to create an article
type ArticleCreatedEvent struct {
	Title              string `json:"title"`
	Description        string `json:"description"`
	AggregateArticleID string `json:"aggregate_article_id"`
}

//ArticleUpdatedEvent Event to update an article
type ArticleUpdatedEvent struct {
	Title              string `json:"title"`
	Description        string `json:"description"`
	AggregateArticleID string `json:"aggregate_article_id"`
}
