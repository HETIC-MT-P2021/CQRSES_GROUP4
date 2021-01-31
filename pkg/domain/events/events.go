package events

// ArticleCreatedEvent Event to create an article
type ArticleCreatedEvent struct {
	ID          int
	Title       string
	Description string
}

// ArticleUpdatedEvent Event to update an article
type ArticleUpdatedEvent struct {
	ID          int
	Title       string
	Description string
}
