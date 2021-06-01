package domain

import (
	"bytes"
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/queries"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/types"
	"github.com/streadway/amqp"
)

func TestInitEventBus(t *testing.T) {
	var cases = []struct {
		what        		string // What I want to test
	}{
		{"Ok"},
	}

	for range cases {
		initEventBus()

		eventBusLength := EventBus.GetLength()
		if (eventBusLength != 2) {
			t.Errorf("eventBusLength = %d, but want %d", eventBusLength, 2)
		}

		eventsName := types.ToSliceByte(EventBus.GetEventsName())
		expectedEventsName := types.ToSliceByte([]string{
			events.ArticleCreatedEventType,
			events.ArticleUpdatedEventType,
		})

		if bytes.Compare(eventsName, expectedEventsName) != 0 {
			t.Errorf("eventsName != expectedEventsName")
		}
	}
}

func TestInitCommandBus(t *testing.T) {
	var cases = []struct {
		what        		string // What I want to test
	}{
		{"Ok"},
	}

	for range cases {
		rabbitImpl := rabbit.NewRabbitRepository(nil, amqp.Queue{})
		initCommandBus(rabbitImpl)

		commandBusLength := CommandBus.GetLength()
		if (commandBusLength != 2) {
			t.Errorf("commandBusLength = %d, but want %d", commandBusLength, 2)
		}

		commandsName := types.ToSliceByte(CommandBus.GetCommandsName())
		expectedCommandsName := types.ToSliceByte([]string{
			pkg.TypeOf(&commands.CreateArticleCommand{}),
			pkg.TypeOf(&commands.UpdateArticleCommand{}),
		})

		if bytes.Compare(commandsName, expectedCommandsName) != 0 {
			t.Errorf("commandsName != expectedCommandsName")
		}
	}
}

func TestInitQueryBus(t *testing.T) {
	var cases = []struct {
		what        		string // What I want to test
	}{
		{"Ok"},
	}

	for range cases {
		initQueryBus()

		queryBusLength := QueryBus.GetLength()
		if (queryBusLength != 1) {
			t.Errorf("queryBusLength = %d, but want %d", queryBusLength, 1)
		}

		queriesName := types.ToSliceByte(QueryBus.GetQueriesName())
		expectedQueriesName := types.ToSliceByte([]string{
			pkg.TypeOf(&queries.ReadArticleQuery{}),
		})

		if bytes.Compare(queriesName, expectedQueriesName) != 0 {
			t.Errorf("queriesName != expectedQueriesName")
		}
	}
}