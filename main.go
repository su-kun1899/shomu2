package main

import (
	"fmt"
	"os"
)

func runCmd(args []string) int {
	fmt.Println("Hello friend!")

	// 引数のチェック
	//if len(args) != 2 {
	//	fmt.Fprintf(os.Stderr, "[ERROR] required two arguments(pattern,package)\n")
	//	return 1
	//}
	//pattern := args[0]
	//packageName := args[1]
	//
	//comments, err := comment.ExtractComments(pattern, packageName)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err)
	//	return 1
	//}
	//
	//for _, c := range comments {
	//	fmt.Printf(c.String())
	//}

	return 0
}

func main() {
	os.Exit(runCmd(os.Args[1:]))
}
