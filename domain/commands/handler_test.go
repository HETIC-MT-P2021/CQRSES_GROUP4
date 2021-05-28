package commands

import (
	"testing"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/cqrs"
	"github.com/stretchr/testify/assert"
)

func getFakeCommandBus() (*cqrs.CommandBus, error) {
	bus := cqrs.NewCommandBus()

	err := bus.AddHandler(
		NewCreateArticleCommandHandler(), 
		&CreateArticleCommand{})

	err = bus.AddHandler(
		NewUpdateArticleCommandHandler(),
		&UpdateArticleCommand{})

	return bus, err
}

func TestCreateArticleCommandHandler(t *testing.T) {
	bus, err := getFakeCommandBus()
	if err != nil {
		t.Errorf(err.Error())
	}

	createArticleCommandOkImpl := cqrs.NewCommandImpl(&CreateArticleCommand{
		Title: "test",
		Description: "test",
	})

	createArticleCommandEmptyTitleImpl := cqrs.NewCommandImpl(&CreateArticleCommand{
		Title: "",
		Description: "test",
	})

	createArticleCommandEmptyDescImpl := cqrs.NewCommandImpl(&CreateArticleCommand{
		Title: "test",
		Description: "",
	})

	var cases = []struct {
		what        							string // What I want to test
		cmdImpl 									cqrs.Command // input
	}{
		{"Ok", createArticleCommandOkImpl},
		{"Empty Title", createArticleCommandEmptyTitleImpl},
		{"Empty Description", createArticleCommandEmptyDescImpl},
	}

	for _, testCase := range cases {
		err := bus.Dispatch(testCase.cmdImpl)
		if testCase.what == "Ok" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestUpdateArticleCommandHandler(t *testing.T) {
	bus, err := getFakeCommandBus()
	if err != nil {
		t.Errorf(err.Error())
	}

	updateArticleCommandOkImpl := cqrs.NewCommandImpl(&UpdateArticleCommand{
		Title: "test",
		Description: "test",
	})

	updateArticleCommandEmptyTitleImpl := cqrs.NewCommandImpl(&UpdateArticleCommand{
		Title: "",
		Description: "test",
	})

	updateArticleCommandEmptyDescImpl := cqrs.NewCommandImpl(&UpdateArticleCommand{
		Title: "test",
		Description: "",
	})

	var cases = []struct {
		what        							string // What I want to test
		cmdImpl 									cqrs.Command // input
	}{
		{"Ok", updateArticleCommandOkImpl},
		{"Empty Title", updateArticleCommandEmptyTitleImpl},
		{"Empty Description", updateArticleCommandEmptyDescImpl},
	}

	for _, testCase := range cases {
		err := bus.Dispatch(testCase.cmdImpl)
		if testCase.what == "Ok" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}