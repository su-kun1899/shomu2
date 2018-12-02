package main

import (
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
	"fmt"
)

func runCmd(args []string) int {
	if len(args) == 0 {
		fmt.Fprintf(os.Stdout, "shomu2 is a tool for trivial things.\n")
		return 0
	}

	commandType := args[0]

	config, err := shomu2.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return shomu2.Fail
	}

	items, err := shomu2.NewItems(config.FileName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return shomu2.Fail
	}

	command, err := shomu2.NewCommand(commandType, items)
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
