package commands

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/events"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/mock"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/rabbit"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	createArticleCommandOk = &CreateArticleCommand{
		Title: "test",
		Description: "test",
	}

	createArticleCommandEmptyTitle = &CreateArticleCommand{
		Title: "",
		Description: "test",
	}

	createArticleCommandEmptyDesc = &CreateArticleCommand{
		Title: "test",
		Description: "",
	}

	updateArticleCommandOk = &UpdateArticleCommand{
		Title: "test",
		Description: "test",
	}

	updateArticleCommandEmptyTitle = &UpdateArticleCommand{
		Title: "",
		Description: "test",
	}

	updateArticleCommandEmptyDesc = &UpdateArticleCommand{
		Title: "test",
		Description: "",
	}
)

type fakeCommandBus struct {
	commandBus *cqrs.CommandBus
	mock *mock.MockRabbitRepository
	err error
}

func getFakeCommandBus(t *testing.T) (*fakeCommandBus) {
	bus := cqrs.NewCommandBus()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mck := mock.NewMockRabbitRepository(ctrl)

	err := bus.AddHandler(
		NewCreateArticleCommandHandler(mck), 
		&CreateArticleCommand{})

	err = bus.AddHandler(
		NewUpdateArticleCommandHandler(mck),
		&UpdateArticleCommand{})

	return &fakeCommandBus {
		commandBus: bus, 
		mock: mck, 
		err: err,
	}
}

func TestCreateArticleCommandHandler(t *testing.T) {
	var cases = []struct {
		what        							string // What I want to test
		cmd       								interface{}
	}{
		{"Ok", createArticleCommandOk},
		{"Empty Title", createArticleCommandEmptyTitle},
		{"Empty Description", createArticleCommandEmptyDesc},
	}

	for _, testCase := range cases {
		fake := getFakeCommandBus(t)
		if fake.err != nil {
			t.Errorf(fake.err.Error())
		}

		fake.mock.
		 	EXPECT().
		 	QueueConnector(rabbit.ConsumeMessage{
				EventType: events.ArticleCreatedEventType,
				Payload: events.ArticleCreatedEvent{
					Title: testCase.cmd.(*CreateArticleCommand).Title,
					Description: testCase.cmd.(*CreateArticleCommand).Description,
				},
			}).
		 	DoAndReturn(func(_ interface{}) error {
		 		return nil
		 	})

		cmdImpl := cqrs.NewCommandImpl(testCase.cmd)

		err := fake.commandBus.Dispatch(cmdImpl)
		if testCase.what == "Ok" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestUpdateArticleCommandHandler(t *testing.T) {
	var cases = []struct {
		what        							string // What I want to test
		cmd 											interface{} // input
	}{
		{"Ok", updateArticleCommandOk},
		{"Empty Title", updateArticleCommandEmptyTitle},
		{"Empty Description", updateArticleCommandEmptyDesc},
	}

	for _, testCase := range cases {
		fake := getFakeCommandBus(t)
		if fake.err != nil {
			t.Errorf(fake.err.Error())
		}

		fake.mock.
		 	EXPECT().
		 	QueueConnector(rabbit.ConsumeMessage{
				EventType: events.ArticleUpdatedEventType,
				Payload: events.ArticleUpdatedEvent{
					Title: testCase.cmd.(*UpdateArticleCommand).Title,
					Description: testCase.cmd.(*UpdateArticleCommand).Description,
				},
			}).
		 	DoAndReturn(func(_ interface{}) error {
		 		return nil
		 	})

		cmdImpl := cqrs.NewCommandImpl(testCase.cmd)

		err := fake.commandBus.Dispatch(cmdImpl)
		if testCase.what == "Ok" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}