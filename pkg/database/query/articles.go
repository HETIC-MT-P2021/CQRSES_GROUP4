package query

const (
	QUERY_CREATE_ARTICLE string = "INSERT INTO meows(titre, content, created_at) VALUES($1, $2, $3)"
)
