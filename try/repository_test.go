package try_test

import (
	"testing"
	"os"
	"github.com/su-kun1899/shomu2/try"
)

func TestNewFileRepository(t *testing.T) {
	t.Run("use file if exists", func(t *testing.T) {
		// given
		fileName := "testdata/shomu2db"

		// when-then: check already exists
		if _, err := os.Stat(fileName); err != nil {
			t.Fatal("unexpected error:", err)
		}

		// when-then
		if _, err := try.NewRepository(fileName); err != nil {
			t.Error("unexpected error:", err)
		}
	})

	t.Run("create new file if not exists", func(t *testing.T) {
		// given
		fileName := os.TempDir() + "new"

		// when-then
		if _, err := try.NewRepository(fileName); err != nil {
			t.Error("unexpected error:", err)
		}

		// when-then
		if _, err := os.Stat(fileName); err != nil {
			t.Error("unexpected error:", err)
		}

		// cleanup
		if err := os.Remove(fileName); err != nil {
			t.Error("unexpected error:", err)
		}
	})
}
