package shomu2

import (
	"os"
	"encoding/base64"
	"fmt"
	"bufio"
)

type Item struct {
	Value string
}

type ItemRepository interface {
	Pop() (*Item, error)
	Push(item *Item) error
}

type FileItemRepository struct {
	fileName string
}

func (r *FileItemRepository) Pop() (*Item, error) {
	fp, err := os.Open(r.fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	items := make([]*Item, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		decoded, err := base64.URLEncoding.DecodeString(scanner.Text())
		if err != nil {
			return nil, err
		}

		items = append(items, &Item{string(decoded)})
	}

	if len(items) == 0 {
		return nil, nil
	}

	// TODO 取り出したItemを消す
	return items[len(items)-1 ], nil
}

func (r *FileItemRepository) Push(item *Item) (error) {
	file, err := os.OpenFile(r.fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	encoded := base64.URLEncoding.EncodeToString([]byte(item.Value))
	_, err = file.WriteString(fmt.Sprintf("%s\n", encoded))

	return err
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(fileName string) (ItemRepository, error) {
	repository := &FileItemRepository{fileName: fileName}
	// TODO テストの移行
	// create file if not exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return repository, err
		}
		defer file.Close()
	}

	return repository, nil
}
