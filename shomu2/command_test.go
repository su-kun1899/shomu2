package shomu2_test

import (
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
	"errors"
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
	callRepository := false
	type fields struct {
		repository shomu2.Repository
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           int
		callRepository bool
		// TODO errorも帰ってきてほしい
	}{
		{
			name: "Less argument error",
			fields: fields{
				repository: nil,
			},
			args: args{
				args: []string{},
			},
			want:           shomu2.Fail,
			callRepository: false,
		},
		{
			name: "Too many arguments error",
			fields: fields{
				repository: nil,
			},
			args: args{
				args: []string{"foo", "bar"},
			},
			want:           shomu2.Fail,
			callRepository: false,
		},
		{
			name: "Repository's error",
			fields: fields{
				repository: &fakeRepository{
					fakeAdd: func(item shomu2.Item) error {
						return errors.New("repository's error")
					},
				},
			},
			args: args{
				args: []string{"foo"},
			},
			want:           shomu2.Fail,
			callRepository: false,
		},
		{
			name: "Pushing item success",
			fields: fields{
				repository: &fakeRepository{
					fakeAdd: func(item shomu2.Item) error {
						callRepository = true
						return nil
					},
				},
			},
			args: args{
				args: []string{"foo"},
			},
			want:           shomu2.Success,
			callRepository: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callRepository = false
			command, err := shomu2.NewCommand("push", tt.fields.repository)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			if got := command.Run(tt.args.args).Code; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push.Run().Code = %v, want %v", got, tt.want)
			}
			if tt.callRepository != callRepository {
				t.Errorf("Push.Run() call Repository = %v, want %v", callRepository, tt.callRepository)
			}
		})
	}
}


type fakeRepository struct {
	shomu2.Repository
	fakeAdd func(item shomu2.Item) error
}

func (repository *fakeRepository) Add(item shomu2.Item) error {
	return repository.fakeAdd(item)
}
