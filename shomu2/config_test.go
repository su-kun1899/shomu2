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
		os.Setenv(shomu2.SHOMU2_HOME_ENV_KEY, home)
		defer os.Unsetenv(shomu2.SHOMU2_HOME_ENV_KEY)

		// when
		got, err := shomu2.NewConfig()

		// then
		if err != nil {
			t.Error("unexpected error:", err)
			return
		}
		if got.Home != home {
			t.Errorf("Home = %v, want %v", got.Home, home)
		}
		if fileName := got.FileName(); fileName != home+".shomu2" {
			t.Errorf("DataFile = %v, want %v", fileName, home+".shomu2")
		}
	})

	t.Run("test configuration", func(t *testing.T) {
		// when
		_, err := shomu2.NewConfig()

		// then
		if err == nil {
			t.Error("expected error did not occur")
		}
	})

}
