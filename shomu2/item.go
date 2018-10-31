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
	items    []*Item
}

func (r *FileItemRepository) load() error {
	fp, err := os.Open(r.fileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	items := make([]*Item, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		decoded, err := base64.URLEncoding.DecodeString(scanner.Text())
		if err != nil {
			return err
		}

		items = append(items, &Item{string(decoded)})
	}
	r.items = items

	return nil
}

func (r *FileItemRepository) save() error {
	file, err := os.OpenFile(r.fileName, os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	s := ""
	for _, it := range r.items {
		encoded := base64.URLEncoding.EncodeToString([]byte(it.Value))
		s += fmt.Sprintf("%s\n", encoded)
	}

	_, err = file.WriteString(s)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileItemRepository) Pop() (*Item, error) {
	if len(r.items) == 0 {
		return nil, nil
	}

	item := r.items[len(r.items)-1 ]
	// 取り出したItemを消す
	r.items = r.items[:max(0, len(r.items)-2)]

	err := r.save()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (r *FileItemRepository) Push(item *Item) (error) {
	r.items = append(r.items, item)
	err := r.save()
	if err != nil {
		return err
	}

	return err
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(fileName string) (ItemRepository, error) {
	repository := &FileItemRepository{fileName: fileName}
	// TODO テストの移行
	// create file if not exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// TODO O_CREATE使えば、意識しなくて良い？
		file, err := os.Create(fileName)
		if err != nil {
			return repository, err
		}
		defer file.Close()
	}
	repository.load()

	return repository, nil
}
