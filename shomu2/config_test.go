package shomu2_test

import (
	"github.com/su-kun1899/shomu2/shomu2"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("test configuration", func(t *testing.T) {
		// given
		home := os.TempDir()
		os.Setenv("SHOMU2_HOME", home)

		// when
		got := shomu2.NewConfig()

		// then
		if got.Home != home {
			t.Errorf("Home = %v, want %v", got.Home, home)
		}
		if fileName := got.FileName(); fileName != home+".shomu2" {
			t.Errorf("DataFile = %v, want %v", fileName, home+".shomu2")
		}
	})
}
