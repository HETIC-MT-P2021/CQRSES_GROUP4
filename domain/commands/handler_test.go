package commands

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
)

func TestCreateArticleCommandHandler(t *testing.T) {
	article := database.Article{
		Title: "test",
		Description: "test",
	}
	emptyArticle := database.Article{}

	var cases = []struct {
		what        							string // What I want to test
		article 									database.Article // input
	}{
		{"Ok", article},
		{"Empty article", emptyArticle},
	}

	for _, testCase := range cases {
		createArticleCommand := cqrs.NewCommandImpl(&CreateArticleCommand{
			Title: testCase.article.Title,
			Description: testCase.article.Description,
		})

		switch cmd := createArticleCommand.Payload().(type) {
		case *CreateArticleCommand:
			if cmd.Title == "" || cmd.Description == "" {
				if testCase.what == "Ok" {
					t.Errorf("Fields should not be empty, title = %s, description = %s", cmd.Title, cmd.Description)
				}
			}

			if events.ArticleCreatedEventType != "ArticleCreatedEvent" {
				t.Errorf("ArticleCreatedEventType = %s, but wanted ArticleCreatedEvent", events.ArticleCreatedEventType)
			}
		default:
			t.Errorf("Bad command type")
		}
	}
}

func TestUpdateArticleCommandHandler(t *testing.T) {
	article := database.Article{
		Title: "test",
		Description: "test",
	}
	emptyArticle := database.Article{}

	var cases = []struct {
		what        							string // What I want to test
		article 									database.Article // input
	}{
		{"Ok", article},
		{"Empty article", emptyArticle},
	}

	for _, testCase := range cases {
		updateArticleCommand := cqrs.NewCommandImpl(&UpdateArticleCommand{
			Title: testCase.article.Title,
			Description: testCase.article.Description,
		})

		switch cmd := updateArticleCommand.Payload().(type) {
		case *UpdateArticleCommand:
			if cmd.Title == "" || cmd.Description == "" {
				if testCase.what == "Ok" {
					t.Errorf("Fields should not be empty, title = %s, description = %s", cmd.Title, cmd.Description)
				}
			}

			if events.ArticleUpdatedEventType != "ArticleUpdatedEvent" {
				t.Errorf("ArticleUpdatedEventType = %s, but wanted ArticleUpdatedEvent", events.ArticleUpdatedEventType)
			}
		default:
			t.Errorf("Bad command type")
		}
	}
}