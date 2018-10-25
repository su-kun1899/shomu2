package shomu2

import "os"

type Item struct {
	Value string
}

type ItemRepository interface {
	Pop() (*Item, error)
}

type FileItemRepository struct {
	fileName string
}

func (*FileItemRepository) Pop() (*Item, error) {
	//TODO "implement me"
	return nil, nil
}

// NewItemRepository is a constructor for ItemRepository
func NewItemRepository(fileName string) (ItemRepository, error) {
	repository := &FileItemRepository{fileName: fileName}
	// TODO テストの移行
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
