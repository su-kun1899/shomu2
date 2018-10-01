package shomu2

import (
	"errors"
	"fmt"
)

type Command interface {
	Run(options ...string)
}

type Push struct {
}

func (*Push) Run(options ...string) {
	panic("implement me")
}

func NewCommand(name string) (Command, error) {
	switch name {
	case "push":
		return &Push{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("command \"%s\" is not exist", name))
	}
}
