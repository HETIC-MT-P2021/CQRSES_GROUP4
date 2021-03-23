package commands

// CreateArticleCommand Command to create an article
type CreateArticleCommand struct {
	Title       string
	Description string
}

// UpdateArticleCommand Command to create an article
type UpdateArticleCommand struct {
	AggregateArticleID string
	Title              string
	Description        string
}

// UpdateArticleTitleCommand Command to update title of an article
type UpdateArticleTitleCommand struct {
	Title              string
	AggregateArticleID string
}
