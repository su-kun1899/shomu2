package try

import (
	"fmt"
	"io/ioutil"
	"os"
	"bufio"
	"strings"
	"encoding/base64"
)

type Repository interface {
	Search(keyword string) ([]Item, error)
	Add(item Item) error
	DeleteAll() error
}

type Item struct {
	Value string
}

type fileRepository struct {
	fileName string
}

func (r *fileRepository) Add(item Item) error {
	file, err := os.OpenFile(r.fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	encoded := base64.URLEncoding.EncodeToString([]byte(item.Value))
	_, err = file.WriteString(fmt.Sprintf("%s\n", encoded))

	return err
}

func (r *fileRepository) DeleteAll() error {
	ioutil.WriteFile(r.fileName, []byte(""), 0666)
	return nil
}

// Search is a function for search items by keyword.
// It returns items which contains keyword in its value.
func (r *fileRepository) Search(keyword string) ([]Item, error) {
	fp, err := os.Open(r.fileName)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	items := make([]Item, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		decoded, err := base64.URLEncoding.DecodeString(scanner.Text())
		if err != nil {
			return items, err
		}

		if line := string(decoded); strings.Contains(line, keyword) {
			items = append(items, Item{line})
		}
	}

	return items, nil
}

// NewRepository is a constructor for Repository
func NewRepository(fileName string) (Repository, error) {
	repository := &fileRepository{fileName: fileName}

	// create file if not exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			return repository, err
		}
		defer file.Close()
	}

	return repository, nil
}
