package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
	"reflect"
)

func TestFileItemRepository(t *testing.T) {
	// given
	fileName := "testdata/shomu2db"
	repository, err := shomu2.NewItemRepository(fileName)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	// when no items
	got, err := repository.Pop()

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	var want *shomu2.Item = nil
	if got != want {
		t.Errorf("FileItemRepository.Pop() = %v, want %v", got, want)
		return
	}

	// TODO when push a item
	item := shomu2.Item{Value: "new item"}
	err = repository.Push(&item)

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	// when pop item
	got, err = repository.Pop()

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	want = &item
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileItemRepository.Pop() = %v, want %v", got, want)
		return
	}
}
