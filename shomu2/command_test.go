package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand(t *testing.T) {
	// when
	command, err := shomu2.NewCommand("add")
	if err != nil {
		t.Error("unexpected error:", err)
	}

	// then
	if _, ok := command.(*shomu2.Push); !ok {
		t.Errorf("command is not push command: %v", command)
	}
}
