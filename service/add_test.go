package service

import (
	"testing"
	"os"
	"bufio"
)

func TestAdd(t *testing.T) {
	type args struct {
		item Item
	}
	tests := []struct {
		name string
		args args
	}{
		{"add item", struct{ item Item }{item: Item{Value: "Hello, world!"}}},
		{"add item", struct{ item Item }{item: Item{Value: "Hello, Alice!"}}},
	}
	defer func() {
		err := removeTestFile()
		if err != nil {
			t.Fatalf("colud not remove test file: %v", err)
		}
	}()

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.item)
			actual, err := readTestFile()
			if err != nil {
				t.Fatalf("colud not read test file. %s", err)
			}
			if actual[index] != tt.args.item.Value {
				t.Fatalf("unexpected item value. %s", actual[index])
			}
		})
	}

	// Append
	// TODO 追記はケースを分けよう
	actual, err := readTestFile()
	if err != nil {
		t.Fatalf("colud not read test file. %s", err)
	}
	if len(actual) != len(tests) {
		t.Fatalf("unexpected item count. %d", len(actual))
	}
}

func testFileName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filename := dir + "/shomu2_test"

	return filename, nil
}

func readTestFile() ([]string, error) {
	fileName, err := testFileName()
	if err != nil {
		return nil, err
	}
	return readFile(fileName)
}

func removeTestFile() error {
	fileName, err := testFileName()
	if err != nil {
		return err
	}
	os.Remove(fileName)

	return nil
}

func readFile(fileName string) ([]string, error) {
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
