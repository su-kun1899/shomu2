package shomu2

import (
	"os"
	"errors"
	"path/filepath"
)

const SHOMU2_HOME_ENV_KEY string = "SHOMU2_HOME"

type Config struct {
	Home string
}

func (c *Config) FileName() string {
	return filepath.Join(c.Home, "data")
}

func NewConfig() (*Config, error) {
	home := os.Getenv(SHOMU2_HOME_ENV_KEY)
	if home == "" {
		return &Config{}, errors.New("shomu2 home is not configured")
	}

	return &Config{home}, nil
}
