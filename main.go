package main

import (
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
)

func runCmd(args []string) int {
	commandType := args[0]

	config := shomu2.NewConfig()
	repository, err := shomu2.NewRepository(config.FileName())
	if err != nil {
		return shomu2.Fail
	}

	command, err := shomu2.NewCommand(commandType, repository)
	if err != nil {
		return shomu2.Fail
	}

	// TODO 可変長引数じゃなくて、配列にしたほうがよさげ
	return command.Run(args[1]).Code
}

func main() {
	os.Exit(runCmd(os.Args[1:]))
}
