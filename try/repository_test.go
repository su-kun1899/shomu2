package try_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/try"
	"strings"
)

func TestNewFileRepository(t *testing.T) {
	t.Run("use file if exists", func(t *testing.T) {
		// given
		fileName := "testdata/shomu2db"

		// when-then: check already exists
		if _, err := os.Stat(fileName); err != nil {
			t.Fatal("unexpected error:", err)
		}

		// when-then
		if _, err := try.NewRepository(fileName); err != nil {
			t.Error("unexpected error:", err)
		}
	})

	t.Run("create new file if not exists", func(t *testing.T) {
		// given
		fileName := os.TempDir() + "new"

		// when-then
		if _, err := try.NewRepository(fileName); err != nil {
			t.Error("unexpected error:", err)
		}

		// when-then
		if _, err := os.Stat(fileName); err != nil {
			t.Error("unexpected error:", err)
		}

		// cleanup
		if err := os.Remove(fileName); err != nil {
			t.Error("unexpected error:", err)
		}
	})
}

func Test_Repository_Add(t *testing.T) {
	// TODO ケース増やして整理したい
	// given
	repository, err := try.NewRepository("testdata/shomu2db")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// when:
	item := try.Item{Value: "HogeHoge"}
	err = repository.Add(item)
	if err != nil {
		t.Error("unexpected error:", err)
	}

	// then:
	items, err := repository.Search(item.Value)
	if size := len(items); size != 1 {
		t.Errorf("Search's result size want %d but got %d", 1, size)
	}
	if actual := items[0]; !reflect.DeepEqual(actual, item) {
		t.Errorf("Search want %v but got %v", actual, item)
	}

	// cleanup
	if err = repository.DeleteAll(); err != nil {
		t.Fatal("unexpected error:", err)
	}
}

func Test_Repository_Search(t *testing.T) {
	// TODO ケース増やして整理したい
	// given
	repository, err := try.NewRepository("testdata/shomu2db")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = repository.Add(try.Item{Value: "Hello, world."})
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// when
	items, err := repository.Search("Hello")
	if err != nil {
		t.Error("unexpected error:", err)
	}

	// then
	if size := len(items); size != 1 {
		t.Errorf("Search's result size want %d but got %d", 1, size)
	}
	for _, item := range items {
		if !strings.Contains(item.Value, "Hello") {
			t.Errorf("Item.value must contain %s but got %s", "Hello", item.Value)
		}
	}

	// cleanup
	if err = repository.DeleteAll(); err != nil {
		t.Fatal("unexpected error:", err)
	}
}
