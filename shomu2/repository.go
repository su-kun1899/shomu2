package shomu2

import (
	"os"
	"fmt"
	"bufio"
)

type Repository struct {
	fileName string
}

func (repo Repository) Add(item Item) error {
	file, err := os.OpenFile(repo.fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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

func (repo Repository) ReadAll() ([]Item, error) {
	fp, err := os.Open(repo.fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	items := make([]Item, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		items = append(items, Item{scanner.Text()})
	}

	return items, nil
}

func NewRepository(filename string) (Repository, error) {
	// ファイルがなかったら作る
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// TODO Untested
		file, err := os.Create(filename)
		if err != nil {
			return Repository{filename}, err
		}
		defer file.Close()
	}

	return Repository{filename}, nil
}
