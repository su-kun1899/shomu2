package shomu2

import "os"

type Config struct {
	Home string
}

func NewConfig() (Config, error) {
	return Config{os.Getenv("SHOMU2_HOME")}, nil
}
