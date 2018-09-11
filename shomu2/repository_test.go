package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
)

func TestAdd(t *testing.T) {
	// given
	fileName := ""
	sut, err := shomu2.NewRepository(fileName)
	if err != nil {
		t.Errorf("NewRepository() error = %v", err)
		return
	}

	// when
	item := shomu2.Item{Value: "Hello, world!"}
	sut.Add(item)

	// and
	actual := sut.ReadAll()

	// then
	if len(actual) != 1 || actual[0] != item {
		t.Errorf("Can't read added item. ReadAll() = %v", actual)
		return
	}
}
