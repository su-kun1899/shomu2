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
	Run(args []string) ExitStatus
}

type Push struct {
	repository Repository
}

// TODO コード値とエラーを返せばそれでよかった？
func (command *Push) Run(args []string) ExitStatus {
	// optionのチェック
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "[ERROR] required 1 arguments(command,item)\n")
		return ExitStatus{Fail}
	}
	itemValue := args[0]

	err := command.repository.Add(Item{Value: itemValue})
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
