package events

// ArticleCreatedEvent Event to create an article
type ArticleCreatedEvent struct {
	ID          string
	Title       string
	Description string
}

// ArticleUpdatedEvent Event to update an article
type ArticleUpdatedEvent struct {
	ID          string
	Title       string
	Description string
}
