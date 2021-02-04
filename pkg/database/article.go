package database

// Article data model on elasticsearch database (index called read-model)
type Article struct {
	ID          string `json:"aggregate_article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
