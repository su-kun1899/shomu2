package shomu2_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/su-kun1899/shomu2/shomu2"
)

func TestNewCommand(t *testing.T) {
	type args struct {
		cmdName    string
		repository shomu2.ItemRepository
	}
	tests := []struct {
		name     string
		args     args
		wantType reflect.Type
	}{
		{
			name:     "create push command",
			args:     args{cmdName: "push", repository: &fakeRepository{}},
			wantType: reflect.TypeOf(&shomu2.Push{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shomu2.NewCommand(tt.args.cmdName, tt.args.repository)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if gotType := reflect.TypeOf(got); gotType != tt.wantType {
				t.Errorf("NewCommand() Type = %v, want %v", gotType, tt.wantType)
			}
		})
	}
}

// TODO パラメータライズ度テストにしたい
func TestNewCommand_Push(t *testing.T) {
	t.Run("not exists command", func(t *testing.T) {
		// when-then
		_, err := shomu2.NewCommand("foo", &fakeRepository{})
		if err == nil {
			t.Error("expected error did not occur")
		}
	})
}

func TestNewCommand_Pop(t *testing.T) {
	t.Run("create pop command", func(t *testing.T) {
		// when
		command, err := shomu2.NewCommand("pop", &fakeRepository{})
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
				data: &fakeRepository{
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
				data: &fakeRepository{
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

type fakeRepository struct {
	shomu2.ItemRepository
	fakePush func(item *shomu2.Item) error
}

func (f *fakeRepository) Push(item *shomu2.Item) error {
	return f.fakePush(item)
}
