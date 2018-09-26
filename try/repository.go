package try

import (
	"os"
	"bufio"
	"strings"
	"encoding/base64"
)

type Repository interface {
	Search(keyword string) ([]Item, error)
}

type Item struct {
	Value string
}

type fileRepository struct {
	fileName string
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
