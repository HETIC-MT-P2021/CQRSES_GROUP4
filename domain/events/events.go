package events

//ArticleCreatedEventType is an event
var ArticleCreatedEventType = "ArticleCreatedEvent"

//ArticleUpdatedEventType is an event
var ArticleUpdatedEventType = "ArticleUpdatedEvent"

//ArticleCreatedEvent Event to create an article
type ArticleCreatedEvent struct {
	ID          string `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//ArticleUpdatedEvent Event to update an article
type ArticleUpdatedEvent struct {
	ID          string `json:"aggregate_article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
