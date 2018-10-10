package shomu2

import "os"

type Config struct {
	Home string
}

func (c *Config) FileName() string {
	return c.Home + "data"
}

func NewConfig() Config {
	// TODO 環境変数が未設定の場合エラー
	return Config{os.Getenv("SHOMU2_HOME")}
}
