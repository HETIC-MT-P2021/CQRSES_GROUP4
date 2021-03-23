package events

//ArticleCreatedEventType is an event
var ArticleCreatedEventType = "ArticleCreatedEvent"

//ArticleUpdatedEventType is an event
var ArticleUpdatedEventType = "ArticleUpdatedEvent"

//ArticleUpdatedTitleEventType is an event
var ArticleUpdatedTitleEventType = "ArticleUpdatedTitleEvent"

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

//ArticleUpdatedTitleEvent Event to update title of an article
type ArticleUpdatedTitleEvent struct {
	Title              string `json:"title"`
	AggregateArticleID string `json:"aggregate_article_id"`
}
