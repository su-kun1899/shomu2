package shomu2

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

type Item struct {
	Value string
}

type Items struct {
	values   []*Item
	fileName string
}

func (items *Items) List() []*Item {
	return items.values
}

func (items *Items) Pop() (*Item, error) {
	if len(items.values) == 0 {
		return nil, nil
	}

	item := items.values[len(items.values)-1 ]
	// 取り出したItemを消す
	items.values = items.values[:max(0, len(items.values)-2)]

	err := items.save()
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (items *Items) Push(item *Item) error {
	items.values = append(items.values, item)
	return items.save()
}

func (items *Items) save() error {
	file, err := os.OpenFile(items.fileName, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	s := ""
	for _, it := range items.values {
		encoded := base64.URLEncoding.EncodeToString([]byte(it.Value))
		s += fmt.Sprintf("%s\n", encoded)
	}

	_, err = file.WriteString(s)
	if err != nil {
		return err
	}

	return nil
}

// NewItems is a constructor for Items
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

	return &Items{values: items, fileName: fileName}, nil
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
