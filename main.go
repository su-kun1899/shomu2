package main

import (
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
	"fmt"
)

func runCmd(args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "shomu2 is a tool for trivial things.\n")
		return 0
	}

	commandType := args[0]

	config := shomu2.NewConfig()
	repository, err := shomu2.NewRepository(config.FileName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return shomu2.Fail
	}

	command, err := shomu2.NewCommand(commandType, repository)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return shomu2.Fail
	}

	code, err := command.Run(args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	return code
}

func main() {
	os.Exit(runCmd(os.Args[1:]))
}
