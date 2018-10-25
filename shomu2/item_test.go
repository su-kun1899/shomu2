package shomu2

import (
	"reflect"
	"testing"
)

func TestFileItemRepository_Pop(t *testing.T) {
	type fields struct {
		fileName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Item
		wantErr bool
	}{
		{
			name:    "No items",
			fields:  fields{fileName: "testdata/shomu2db"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository, err := NewItemRepository(tt.fields.fileName)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			got, err := repository.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("FileItemRepository.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileItemRepository.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
