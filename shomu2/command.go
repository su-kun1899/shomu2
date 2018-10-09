package shomu2

import (
	"errors"
	"fmt"
)

const (
	Success = 0
	Fail    = 1
)

type ExitStatus struct {
	Code int
}

type Command interface {
	Run(options ...string) ExitStatus
}

type Push struct {
	repository Repository
}

func (command *Push) Run(options ...string) ExitStatus {
	err := command.repository.Add(Item{Value: options[0]})
	if err != nil {
		// TODO エラーのステータスを返す
	}
	return ExitStatus{Success}
}

func NewCommand(name string, repository Repository) (Command, error) {
	switch name {
	case "push":
		return &Push{repository: repository}, nil
	default:
		return nil, errors.New(fmt.Sprintf("command \"%s\" is not exist", name))
	}
}
