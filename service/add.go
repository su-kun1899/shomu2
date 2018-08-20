package service

import (
	"os"
	"fmt"
)

type Item struct {
	Value string
}

func Add(item Item) error {
	// TODO ファイル名を環境変数で管理したい
	filename := "shomu2_test"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// TODO メソッドにしたい
		// TODO なかったら作るのは最初にやってしまう？
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n", item.Value))
	if err != nil {
		return err
	}

	return nil
}
