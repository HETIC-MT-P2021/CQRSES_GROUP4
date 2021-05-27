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

func TestArticleCreatedEventHandler(t *testing.T) {
	bus, err := getFakeEventBus()
	if err != nil {
		t.Errorf(err.Error())
	}

	articleCreatedEventOKImpl := event.NewEventImpl(ArticleCreatedEventType, &ArticleCreatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	articleCreatedEventBadImpl := event.NewEventImpl("fail", &ArticleCreatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	var cases = []struct {
		what        							string // What I want to test
		eventImpl									event.Event // input
	}{
		{"Ok", articleCreatedEventOKImpl},
		{"ArticleCreatedEventBadImpl Bad", articleCreatedEventBadImpl},
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

func TestArticleUpdatedEventHandler(t *testing.T) {
	bus, err := getFakeEventBus()
	if err != nil {
		t.Errorf(err.Error())
	}

	articleUpdatedEventOKImpl := event.NewEventImpl(ArticleUpdatedEventType, &ArticleUpdatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	articleUpdatedEventBadImpl := event.NewEventImpl("fail", &ArticleUpdatedEvent{
		Title: "title",
		Description: "description",
	}, true)

	var cases = []struct {
		what        							string // What I want to test
		eventImpl									event.Event // input
	}{
		{"Ok", articleUpdatedEventOKImpl},
		{"ArticleUpdatedEventBadImpl Bad", articleUpdatedEventBadImpl},
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