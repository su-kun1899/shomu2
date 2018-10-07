package shomu2

import (
	"errors"
	"fmt"
)

type Command interface {
	Run(options ...string)
}

type Push struct {
	repository Repository
}

func (command *Push) Run(options ...string) {
	command.repository.Add(Item{Value: options[0]})
}

func NewCommand(name string, repository Repository) (Command, error) {
	switch name {
	case "push":
		return &Push{repository: repository}, nil
	default:
		return nil, errors.New(fmt.Sprintf("command \"%s\" is not exist", name))
	}
}
