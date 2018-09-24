package shomu2_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
)

func TestAdd(t *testing.T) {
	// given
	fileName := os.TempDir() + "test.shomu2"
	sut, err := shomu2.NewRepository(fileName)
	if err != nil {
		t.Errorf("NewRepository() error = %v", err)
		return
	}

	// when
	item := shomu2.Item{Value: "Hello, world!"}
	err = sut.Add(item)
	if err != nil {
		t.Errorf("Add() error = %v", err)
	}

	// and
	actual, err := sut.ReadAll()
	if err != nil {
		t.Errorf("ReadAll() error = %v", err)
	}

	// then
	if len(actual) != 1 || actual[0] != item {
		t.Errorf("Can't read added item. ReadAll() = %v, sut = %v", actual, sut)
		return
	}

	// cleanup
	os.Remove(fileName)
}

func TestNewRepository(t *testing.T) {
	newFileName := os.TempDir() + "new"
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    shomu2.Repository
		wantErr bool
	}{
		{"empty file name", args{""}, shomu2.Repository{}, true},
		{"new file name", args{newFileName}, shomu2.Repository{FileName: newFileName}, false},
		{"exists file name", args{"testdata/exists"}, shomu2.Repository{FileName: "testdata/exists"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shomu2.NewRepository(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
