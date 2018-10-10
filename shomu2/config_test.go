package shomu2

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("test configuration", func(t *testing.T) {
		// given
		home := os.TempDir()
		os.Setenv("SHOMU2_HOME", home)

		// when
		got, err := NewConfig()

		// then
		if err != nil {
			t.Error("unexpected error:", err)
			return
		}
		if got.Home != home {
			t.Errorf("Home = %v, want %v", got, home)
		}
	})
}
