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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Add(tt.args.item)
			actual, err := readTestFile()
			if err != nil {
				t.Fatalf("colud not read test file. %s", err)
			}
			if len(actual) != 1 {
				t.Fatalf("unexpected item count. %d", len(actual))
			}
			if actual[0] != tt.args.item.Value {
				t.Fatalf("unexpected item value. %s", actual[0])
			}
		})
	}
}

func readTestFile() ([]string, error) {
	dir, _ := os.Getwd()
	filename := dir + "/shomu2_test"

	return readFile(filename)
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
