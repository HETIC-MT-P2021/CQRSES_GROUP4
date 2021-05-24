package events

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
)

type inputGetPayloadMapped struct {
	expectedTitle string
	expectedDescription string
	expectedID string
}

func TestGetPayloadMapped(t *testing.T) {
	inputGetPayloadMappedOK := inputGetPayloadMapped{
		expectedTitle: "example",
		expectedDescription: "example",
		expectedID: "1",
	}

	inputGetPayloadMappedEmpty := inputGetPayloadMapped{}

	var cases = []struct {
		what        							string // What I want to test
		input 									inputGetPayloadMapped // input
	}{
		{"Ok", inputGetPayloadMappedOK},
		{"Empty input", inputGetPayloadMappedEmpty},
	}
	
	for _, testCase := range cases {
		input := testCase.input
		evPayload := ArticleCreatedEvent{
			Title: input.expectedTitle,
			Description: input.expectedDescription,
			AggregateArticleID: input.expectedID,
		}
		ev := event.NewEventImpl(ArticleCreatedEventType,
			evPayload, true)

		payloadMapped, err := getPayloadMapped(ev)
		if err != nil {
			t.Errorf("error = %s", err)
		}
		
		if payloadMapped["title"] != input.expectedTitle {
			t.Errorf("payloadMapped['title'] = %s, want %s", payloadMapped["title"], input.expectedTitle)
		}

		if payloadMapped["description"] != input.expectedDescription {
			t.Errorf("payloadMapped['description'] = %s, want %s", payloadMapped["description"], input.expectedDescription)
		}

		if payloadMapped["aggregate_article_id"] != input.expectedID {
			t.Errorf("payloadMapped['aggregate_article_id'] = %s, want %s", payloadMapped["aggregate_article_id"], input.expectedID)
		}
	}
}