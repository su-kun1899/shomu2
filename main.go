package main

import (
	"fmt"
	"os"
)

// exitStatus defines statuses returned when command exit
type exitStatus = int

// Exit Statuses
const (
	success            exitStatus = 0
	lackOfRequiredArgs exitStatus = 11
	notExistCommand    exitStatus = 12
)

func runCmd(args []string) exitStatus {
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
