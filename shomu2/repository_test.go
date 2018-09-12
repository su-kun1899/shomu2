package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
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
