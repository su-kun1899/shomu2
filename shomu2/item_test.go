package shomu2

import (
	"reflect"
	"testing"
)

func TestFileItemRepository(t *testing.T) {
	// given
	fileName := "testdata/shomu2db"
	repository, err := NewItemRepository(fileName)
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
	if !reflect.DeepEqual(got, nil) {
		t.Errorf("FileItemRepository.Pop() = %v, want %v", got, tt.want)
	}

	// TODO when push a item
}
