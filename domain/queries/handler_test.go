package queries

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
)

func TestReadArticleQueryHandler(t *testing.T) {
	var cases = []struct {
		what        							string // What I want to test
		aggregateArticleID 				string // input
	}{
		{"Ok", "1"},
		{"aggregateArticleID empty", ""},
	}

	for _, testCase := range cases {
		readArticleQuery := cqrs.NewQueryImpl(&ReadArticleQuery{
			AggregateArticleID: testCase.aggregateArticleID,
		})

		switch readArticleQuery.Payload().(type) {
		case *ReadArticleQuery:
			aggregateArticleID := readArticleQuery.Payload().(*ReadArticleQuery).AggregateArticleID

			if aggregateArticleID == "" && testCase.what == "Ok" {
				t.Errorf("aggregateArticleID should not be empty")
			}
		default:
			t.Errorf("Bad command type")
		}
	}
}