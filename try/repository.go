package try

import "os"

type Repository interface {
	findAll() ([]Item, error)
}

type Item struct {
	Value string
}

type fileRepository struct {
	fileName string
}

func (r *fileRepository) findAll() ([]Item, error) {
	panic("implement me")
}

// NewRepository is a constructor for Repository
func NewRepository(fileName string) (Repository, error) {
	repository := new(fileRepository)
	repository.fileName = fileName

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
