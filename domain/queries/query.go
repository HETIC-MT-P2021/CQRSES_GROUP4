package queries

// ReadArticleQuery Query to get one article
type ReadArticleQuery struct {
	AggregateArticleID string `json:"aggregate_article_id"`
}
