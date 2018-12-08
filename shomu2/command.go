package shomu2

import (
	"errors"
	"fmt"
	"os"
)

const (
	Success = 0
	Fail    = 1
)

type Command interface {
	Run(args []string) (int, error)
}

type Push struct {
	data Data
}

type Data interface {
	Push(item *Item) error
	Pop() (*Item, error)
}

func (command *Push) Run(args []string) (int, error) {
	// optionのチェック
	if len(args) != 1 {
		return Fail, errors.New(fmt.Sprintf("called by illegal arguments: %v", args))
	}
	itemValue := args[0]

	err := command.data.Push(&Item{Value: itemValue})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return Fail, err
	}
	return Success, nil
}

func NewCommand(name string, data Data) (Command, error) {
	switch name {
	case "push":
		return &Push{data: data}, nil
	default:
		return nil, errors.New(fmt.Sprintf("command \"%s\" is not exist", name))
	}
}
