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

type Items struct {
	Values []*Item
}

func NewItems(fileName string) (*Items, error) {
	// load items
	fp, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
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

	return &Items{Values: items}, nil
}

// TODO Repositoryにする
type FileRepository interface {
	FindAll() ([]*Item, error)
	SaveAll([]*Item) error
}

type FileItems struct {
	values []*Item
	// TODO ItemRepositoryにする
	repo FileRepository
}

func (items *FileItems) Pop() (*Item, error) {
	panic("implement me")
}

func (*FileItems) Push(item *Item) error {
	panic("implement me")
}

type ItemRepository interface {
	Pop() (*Item, error)
	Push(item *Item) error
}

type FileItemRepository struct {
	fileName string
	items    []*Item
}

func (r *FileItemRepository) save() error {
	file, err := os.OpenFile(r.fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
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
	return r.save()
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(fileName string) (ItemRepository, error) {
	// load items
	fp, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
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

	return &FileItemRepository{fileName: fileName, items: items}, nil
}
