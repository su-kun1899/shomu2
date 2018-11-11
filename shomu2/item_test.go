package shomu2_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
)

func TestFileItemRepository(t *testing.T) {
	// TODO このケースはさよならかな。。
	t.Skip()

	// given
	fileName := filepath.Join(os.TempDir(), "items")
	defer func() {
		if err := os.Remove(fileName); err != nil {
			t.Error("unexpected error:", err)
			return
		}
	}()

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

	// when push a item
	item := shomu2.Item{Value: "new item"}
	err = repository.Push(&item)

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	// and pop item
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

func TestNewItems(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []*shomu2.Item
		wantErr bool
	}{
		{
			name:    "Empty file",
			args:    args{fileName: filepath.Join("testdata", "empty")},
			want:    []*shomu2.Item{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shomu2.NewItems(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if values := got.Values; !reflect.DeepEqual(values, tt.want) {
				t.Errorf("NewItems() = %v, want %v", values, tt.want)
			}
		})
	}
}
