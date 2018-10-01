package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand_push(t *testing.T) {
	t.Run("create push command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("push")
		if err != nil {
			t.Error("unexpected error:", err)
		}

		// then
		if _, ok := command.(*shomu2.Push); !ok {
			t.Errorf("command is not push command: %v", command)
		}
	})

	t.Run("not exists command", func(t *testing.T) {
		// when-then
		_, err := shomu2.NewCommand("foo")
		if err == nil {
			t.Error("expected error did not occur")
		}
	})
}
