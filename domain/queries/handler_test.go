package queries

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	article = database.Article{
		ID: "1",
		Title: "test",
		Description: "test",
	}
)

const (
	VALID_AGGREGATE_ARTICLE_ID = "1"
)

func TestReadArticleQueryHandler(t *testing.T) {
	var cases = []struct {
		what        							string // What I want to test
		aggregateArticleID				string // input
	}{
		{"Ok", VALID_AGGREGATE_ARTICLE_ID},
	}

	for _, testCase := range cases {
		bus := cqrs.NewQueryBus()

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mck := mock.NewMockRepository(ctrl)

		queryHandler := NewReadArticleQueryHandler(mck)
		err := bus.AddHandler(queryHandler, &ReadArticleQuery{})
		if err != nil {
		}

		mck.
			EXPECT().
			GetReadmodel(testCase.aggregateArticleID).
			DoAndReturn(func(_ string) (database.Article, error) {
				return article, nil
			})

		cmdImpl := cqrs.NewCommandImpl(&ReadArticleQuery{
			AggregateArticleID: testCase.aggregateArticleID,
		})

		_, err = bus.Dispatch(cmdImpl)
		
		assert.NoError(t, err)
	}
}