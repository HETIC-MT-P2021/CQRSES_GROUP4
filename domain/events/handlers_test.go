package events

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/event"
	"github.com/stretchr/testify/assert"
)

func getFakeEventBus() (*event.EventBus, error) {
	bus := event.NewEventBus()

	err := bus.AddHandler(
		NewArticleCreatedEventHandler(), 
		ArticleCreatedEventType)

	err = bus.AddHandler(
		NewArticleUpdatedEventHandler(),
		ArticleUpdatedEventType)

	return bus, err
}

func TestAllEvents(t *testing.T) {
	bus, err := getFakeEventBus()
	if err != nil {
		t.Errorf(err.Error())
	}

	ArticleCreatedEventOKImpl := event.NewEventImpl(ArticleCreatedEventType, &ArticleCreatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	ArticleUpdatedEventOKImpl := event.NewEventImpl(ArticleUpdatedEventType, &ArticleUpdatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	ArticleCreatedEventBadImpl := event.NewEventImpl("fail", &ArticleCreatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	ArticleUpdatedEventBadImpl := event.NewEventImpl("fail", &ArticleUpdatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	var cases = []struct {
		what        							string // What I want to test
		eventImpl									event.Event // input
	}{
		{"Ok", ArticleCreatedEventOKImpl},
		{"Ok", ArticleUpdatedEventOKImpl},
		{"ArticleCreatedEventBadImpl Bad", ArticleCreatedEventBadImpl},
		{"ArticleUpdatedEventBadImpl Bad", ArticleUpdatedEventBadImpl},
	}

	for _, testCase := range cases {
		err := bus.Dispatch(testCase.eventImpl)
		if testCase.what == "Ok" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}