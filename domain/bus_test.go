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

		eventsName, err := types.StringToSliceByte(EventBus.GetEventsName())
		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}
		expectedEventsName, err := types.StringToSliceByte([]string{
			events.ArticleCreatedEventType,
			events.ArticleUpdatedEventType,
		})

		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}

		if !bytes.Equal(eventsName, expectedEventsName) {
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

		commandsName, err := types.StringToSliceByte(CommandBus.GetCommandsName())
		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}
		expectedCommandsName, err := types.StringToSliceByte([]string{
			pkg.TypeOf(&commands.CreateArticleCommand{}),
			pkg.TypeOf(&commands.UpdateArticleCommand{}),
		})

		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}

		if !bytes.Equal(commandsName, expectedCommandsName) {
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

		queriesName, err := types.StringToSliceByte(QueryBus.GetQueriesName())
		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}

		expectedQueriesName, err := types.StringToSliceByte([]string{
			pkg.TypeOf(&queries.ReadArticleQuery{}),
		})
		if err != nil {
			t.Errorf("got err : %s", err.Error())
		}

		if !bytes.Equal(queriesName, expectedQueriesName) {
			t.Errorf("queriesName != expectedQueriesName")
		}
	}
}