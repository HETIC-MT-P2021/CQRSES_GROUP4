package cqrs

import (
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/domain/commands"
	"testing"
)

func TestCommandBus_Dispatch(t *testing.T) {
	command := NewCommandImpl(&commands.UpdateArticleCommand{
		AggregateArticleID: "3242",
		Title:              "test",
		Description:        "Lorem ipsum dolor sit amet.",
	})

	fmt.Println(command.Payload())
}
