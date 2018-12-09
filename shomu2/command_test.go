package shomu2_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand_Push(t *testing.T) {
	t.Run("create push command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("push", &fakeData{})
		if err != nil {
			t.Fatal("unexpected error:", err)
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

func TestNewCommand_Pop(t *testing.T) {
	t.Run("create pop command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("pop", &fakeData{})
		if err != nil {
			t.Fatal("unexpected error:", err)
		}

		// then
		if _, ok := command.(*shomu2.Pop); !ok {
			t.Errorf("command is %v, want %s", reflect.ValueOf(command).Type(), "*shomu2.Pop")
		}
	})
}

func TestPush_Run(t *testing.T) {
	called := false
	type fields struct {
		data shomu2.ItemRepository
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		called  bool
		want    int
		wantErr bool
	}{
		{
			name: "Less argument error",
			fields: fields{
				data: nil,
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
				data: nil,
			},
			args: args{
				args: []string{"foo", "bar"},
			},
			want:    shomu2.Fail,
			wantErr: true,
		},
		{
			name: "Command error",
			fields: fields{
				data: &fakeData{
					fakePush: func(item *shomu2.Item) error {
						return errors.New("command error")
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
				data: &fakeData{
					fakePush: func(item *shomu2.Item) error {
						called = true
						return nil
					},
				},
			},
			args: args{
				args: []string{"foo"},
			},
			called: true,
			want:   shomu2.Success,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			called = false
			command, err := shomu2.NewCommand("push", tt.fields.data)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}
			got, err := command.Run(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push.Run() = %v, want %v", got, tt.want)
			}
			if tt.called != called {
				t.Errorf("Push.Run() call Repository = %v, want %v", called, tt.called)
			}
		})
	}
}

type fakeData struct {
	shomu2.ItemRepository
	fakePush func(item *shomu2.Item) error
}

func (f *fakeData) Push(item *shomu2.Item) error {
	return f.fakePush(item)
}
