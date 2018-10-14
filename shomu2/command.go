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

type ExitStatus struct {
	Code int
}

type Command interface {
	// TODO 消す
	Run1(args []string) ExitStatus
	Run(args []string) (int, error)
}

type Push struct {
	repository Repository
}

func (command *Push) Run(args []string) (int, error) {
	// optionのチェック
	if len(args) != 1 {
		return Fail, errors.New(fmt.Sprintf("called by illegal arguments: %v", args))
	}
	itemValue := args[0]

	err := command.repository.Add(Item{Value: itemValue})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return Fail, err
	}
	return Success, nil
}

// TODO 消す
func (command *Push) Run1(args []string) ExitStatus {
	// optionのチェック
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "[ERROR] required 1 arguments(command,item)\n")
		return ExitStatus{Fail}
	}
	itemValue := args[0]

	err := command.repository.Add(Item{Value: itemValue})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return ExitStatus{Fail}
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
