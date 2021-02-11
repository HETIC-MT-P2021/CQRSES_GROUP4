package commands

// CreateArticleCommand Command to create an article
type CreateArticleCommand struct {
	Title       string
	Description string
}

// UpdateArticleCommand Command to create an article
type UpdateArticleCommand struct {
	ID          string
	Title       string
	Description string
}
