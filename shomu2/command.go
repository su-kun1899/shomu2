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

type ItemRepository interface {
	Push(item *Item) error
	Pop() (*Item, error)
}

type Push struct {
	repository ItemRepository
}

func (command *Push) Run(args []string) (int, error) {
	// optionのチェック
	if len(args) != 1 {
		return Fail, errors.New(fmt.Sprintf("called by illegal arguments: %v", args))
	}
	itemValue := args[0]

	err := command.repository.Push(&Item{Value: itemValue})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return Fail, err
	}
	return Success, nil
}

type Pop struct {
	repository ItemRepository
}

func (*Pop) Run(args []string) (int, error) {
	panic("implement me")
}

// TODO 引数名Rename
func NewCommand(name string, data ItemRepository) (Command, error) {
	switch name {
	case "push":
		return &Push{repository: data}, nil
	case "pop":
		return &Pop{repository: data}, nil
	default:
		return nil, errors.New(fmt.Sprintf("command \"%s\" is not exist", name))
	}
}
