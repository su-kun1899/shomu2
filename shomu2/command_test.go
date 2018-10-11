package shomu2_test

import (
	"testing"
	"github.com/su-kun1899/shomu2/shomu2"
	"reflect"
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
	param := shomu2.Item{Value: "push!push!"}

	// and: mocking repository
	called := false
	repository := fakeRepository{fakeAdd: func(item shomu2.Item) error {
		called = reflect.DeepEqual(param, item)
		return nil
	}}

	// and: create command
	command, err := shomu2.NewCommand("push", &repository)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	// when
	exitStatus := command.Run([]string{param.Value})

	// then
	if exitStatus.Code != shomu2.Success {
		t.Errorf("Run want %v but got %v", shomu2.Success, exitStatus.Code)
	}
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
