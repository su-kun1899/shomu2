package shomu2_test

import (
	"reflect"
	"testing"

	"errors"
	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand_push(t *testing.T) {
	t.Run("create push command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("push", &fakeData{})
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
		_, err := shomu2.NewCommand("foo", &fakeData{})
		if err == nil {
			t.Error("expected error did not occur")
		}
	})
}

func TestPush_Run(t *testing.T) {
	// TODO 書き換え
	t.Skip()
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
		callRepository bool
		want           int
		wantErr        bool
	}{
		{
			name: "Less argument error",
			fields: fields{
				repository: nil,
			},
			args: args{
				args: []string{},
			},
			want:    shomu2.Fail,
			wantErr: true,
		},
		{
			name: "Too many arguments error",
			fields: fields{
				repository: nil,
			},
			args: args{
				args: []string{"foo", "bar"},
			},
			want:    shomu2.Fail,
			wantErr: true,
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
			want:    shomu2.Fail,
			wantErr: true,
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
			callRepository: true,
			want:           shomu2.Success,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callRepository = false
			command, err := shomu2.NewCommand("push", nil)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			got, err := command.Run(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push.Run() = %v, want %v", got, tt.want)
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

type fakeData struct {
	shomu2.Data
	fakePush func(item *shomu2.Item) error
}

func (f *fakeData) Push(item *shomu2.Item) error {
	return f.fakePush(item)
}
