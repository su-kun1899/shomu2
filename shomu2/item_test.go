package shomu2_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
)

func TestItems(t *testing.T) {
	// given
	fileName := filepath.Join(os.TempDir(), "items")
	defer os.Remove(fileName)

	items, err := shomu2.NewItems(fileName)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	// when no items
	got, err := items.Pop()

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
	err = items.Push(&item)

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}

	// and pop item
	got, err = items.Pop()

	// then
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	want = &item
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Items.Pop() = %v, want %v", got, want)
		return
	}
}

func TestNewItems(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "TestNewItems")
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	defer os.RemoveAll(tempDir)

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
		{
			name:    "New file",
			args:    args{fileName: filepath.Join(tempDir, "tmpfile")},
			want:    []*shomu2.Item{},
			wantErr: false,
		},
		{
			name: "Existed file",
			args: args{fileName: filepath.Join("testdata", "data")},
			want: []*shomu2.Item{
				{Value: "Hello!"},
				{Value: "My name is shomu2."},
				{Value: "Nice to meet you."},
				{Value: "Good-bye!"},
			},
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
			if values := got.List(); !equal(t, values, tt.want) {
				t.Errorf("NewItems() = %v, want %v", values, tt.want)
			}
		})
	}
}

func equal(t *testing.T, values1, values2 []*shomu2.Item) bool {
	t.Helper()

	for i := range values1 {
		if *values1[i] != *values2[i] {
			return false
		}
	}

	return true
}
