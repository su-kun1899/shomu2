package service

import (
	"os"
	"fmt"
)

type Item struct {
	Value string
}

func Add(item Item) error {
	println("item value: ", item.Value)
	// TODO ファイル名を環境変数で管理したい
	file, err := os.Create("shomu2_test")
	if err != nil {
		return err
	}
	fmt.Printf("file name: %s\n", file.Name())

	return nil
}
