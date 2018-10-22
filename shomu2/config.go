package shomu2

import (
	"os"
	"errors"
)

type Config struct {
	Home string
}

func (c *Config) FileName() string {
	return c.Home + ".shomu2"
}

func NewConfig() (*Config, error) {
	home := os.Getenv("SHOMU2_HOME")
	if home == "" {
		return &Config{}, errors.New("shomu2 home is not configured")
	}

	return &Config{home}, nil
}
