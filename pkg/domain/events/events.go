package events

type ArticleCreatedEvent struct {
	ID          int
	Title       string
	Description string
}

type ArticleUpdatedEvent struct {
	ID          int
	Title       string
	Description string
}
