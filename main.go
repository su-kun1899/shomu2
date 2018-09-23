package main

import (
	"fmt"
	"os"
)

func runCmd(args []string) int {
	// 引数のチェック
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "[ERROR] required two arguments(command,item)\n")
		return 1
	}

	command := args[0]
	item := args[1]
	switch command {
	case "add":
		// TODO
		fmt.Fprintf(os.Stdout, "item: %v を追加するよ\n", item)
	default:
		fmt.Fprintf(os.Stderr, "[ERROR] not exist command: %s \n", command)
		return 1
	}

	return 0
}

func main() {
	os.Exit(runCmd(os.Args[1:]))
}
