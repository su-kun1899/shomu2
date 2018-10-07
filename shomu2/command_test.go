package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand_push(t *testing.T) {
	t.Run("create push command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("push", &fakeRepository{})
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
		_, err := shomu2.NewCommand("foo", &fakeRepository{})
		if err == nil {
			t.Error("expected error did not occur")
		}
	})
}

func TestPush_Run(t *testing.T) {
	// given
	called := false
	repository := fakeRepository{fakeAdd: func(item shomu2.Item) error {
		called = true
		return nil
	}}
	command, err := shomu2.NewCommand("push", &repository)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// and
	item := shomu2.Item{Value: "push!push!"}

	// when
	command.Run(item.Value)

	// then
	if !called {
		t.Error("Repository is not called")
	}
}

type fakeRepository struct {
	shomu2.Repository
	fakeAdd func(item shomu2.Item) error
}

func (repository *fakeRepository) Add(item shomu2.Item) error {
	return repository.fakeAdd(item)
}
