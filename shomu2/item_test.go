package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
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
	if got != nil {
		t.Errorf("FileItemRepository.Pop() = %v, want %v", got, nil)
	}

	// TODO when push a item
}
