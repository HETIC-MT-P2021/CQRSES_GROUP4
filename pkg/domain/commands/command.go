package commands

import (
	db "github.com/jibe0123/CQRSES_GROUP4/pkg/database"
)

// Command Tells to system to make modifications
// Type can be AddArticle or EditArticle
type Command struct {
	Type    string
	Payload interface{}
}

type CreateArticleCommand struct {
	Article db.Article
}

type UpdateArticleCommand struct {
	Article   db.Article
	ArticleID int
}
