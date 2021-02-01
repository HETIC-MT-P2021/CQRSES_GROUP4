package commands

// CreateArticleCommand Command to create an article
type CreateArticleCommand struct {
	ID          int
	Title       string
	Description string
}

// UpdateArticleCommand Command to create an article
type UpdateArticleCommand struct {
	ID          int
	Title       string
	Description string
}
