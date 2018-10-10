package shomu2

import "os"

type Config struct {
	Home string
}

func (c *Config) FileName() string {
	return c.Home + "data"
}

func NewConfig() Config {
	return Config{os.Getenv("SHOMU2_HOME")}
}
