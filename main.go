package main

import (
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
	"path/filepath"
)

// TODO こいつ自体も場所変えていいのかも。。
func runCmd(args []string) shomu2.ExitStatus {
	commandType := args[0]
	// TODO ファイルの場所をユーザーディレクトリとかにする？
	homeDir := os.Getenv("SHOMU2_HOME")
	if homeDir == "" {
		// TODO
		panic("SHOMU2_HOMEが未設定")
	}

	repository, err := shomu2.NewRepository(homeDir + filepath.FromSlash("/data"))
	if err != nil {
		// TODO
		panic(err)
	}

	command, err := shomu2.NewCommand(commandType, repository)
	if err != nil {
		// TODO
		panic(err)
	}

	// TODO 可変長引数じゃなくて、配列にしたほうがよさげ
	return command.Run(args[1])
}

func main() {
	status := runCmd(os.Args[1:])
	os.Exit(status.Code)
}
