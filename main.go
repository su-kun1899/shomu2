package main

import (
	"fmt"
	"os"
)

// ExitStatus defines statuses returned when command exit
type ExitStatus = int

// Exit Statuses
const (
	success            ExitStatus = 0
	lackOfRequiredArgs ExitStatus = 11
	notExistCommand    ExitStatus = 12
)

func runCmd(args []string) ExitStatus {
	// 引数のチェック
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "[ERROR] required two arguments(command,item)\n")
		return lackOfRequiredArgs
	}

	command := args[0]
	item := args[1]
	switch command {
	case "add":
		// TODO
		fmt.Fprintf(os.Stdout, "item: %v を追加するよ\n", item)
	default:
		fmt.Fprintf(os.Stderr, "[ERROR] not exist command: %s \n", command)
		return notExistCommand
	}

	return success
}

func main() {
	os.Exit(runCmd(os.Args[1:]))
}
