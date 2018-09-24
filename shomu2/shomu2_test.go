package shomu2

import (
	"testing"
	"bytes"
	"io"
)

// TODO productionコードにする
// TODO rename
type Repo interface {
	Save(item Item) error
}

type BuffRepo struct {
	Buff *bytes.Buffer
}

func (r *BuffRepo) Save(item Item) error {
	_, err := r.Buff.Write([]byte(item.Value))
	return err
}

func (r *BuffRepo) FindAll() ([]Item, error) {
	items := make([]Item, 0)
	for {
		line, err := r.Buff.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return items, err
		}
		items = append(items, Item{string(line)})
	}
	return items, nil
}

type Shomu2 struct {
	Repo Repo
}

func (s *Shomu2) Add(item Item) error {
	return s.Repo.Save(item)
}

func TestShomu2_Add(t *testing.T) {
	cases := []struct {
		name string
		item Item
	}{
		{
			name: "add item",
			item: Item{Value: "Hello, world"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sut := Shomu2{&BuffRepo{Buff: new(bytes.Buffer)}}
			sut.Add(c.item)

			if buff, ok := sut.Repo.(*BuffRepo); ok {
				actual := buff.Buff.String()
				if actual != c.item.Value {
					t.Errorf("item value want %s but got %s", c.item.Value, actual)
				}
			}
		})
	}
}
