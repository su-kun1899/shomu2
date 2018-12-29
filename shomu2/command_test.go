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
		wantErr  bool
	}{
		{
			name:     "create push command",
			args:     args{cmdName: "push", repository: &fakeRepository{}},
			wantType: reflect.TypeOf(&shomu2.Push{}),
		},
		{
			name:     "create pop command",
			args:     args{cmdName: "pop", repository: &fakeRepository{}},
			wantType: reflect.TypeOf(&shomu2.Pop{}),
		},
		{
			name:     "not exists command",
			args:     args{cmdName: "dummy", repository: &fakeRepository{}},
			wantType: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shomu2.NewCommand(tt.args.cmdName, tt.args.repository)
			if !tt.wantErr && err != nil {
				t.Fatal("unexpected error:", err)
			}

			if gotType := reflect.TypeOf(got); gotType != tt.wantType {
				t.Errorf("NewCommand() Type = %v, want %v", gotType, tt.wantType)
			}
		})
	}
}

func TestPush_Run(t *testing.T) {
	called := false
	type fields struct {
		repo shomu2.ItemRepository
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		called bool
		want   int
	}{
		{
			name: "Less argument error",
			fields: fields{
				repo: nil,
			},
			args: args{
				args: []string{},
			},
			want: shomu2.Fail,
		},
		{
			name: "Too many arguments error",
			fields: fields{
				repo: nil,
			},
			args: args{
				args: []string{"foo", "bar"},
			},
			want: shomu2.Fail,
		},
		{
			name: "Command error",
			fields: fields{
				repo: &fakeRepository{
					fakePush: func(item *shomu2.Item) error {
						return errors.New("command error")
					},
				},
			},
			args: args{
				args: []string{"foo"},
			},
			want: shomu2.Fail,
		},
		{
			name: "Pushing item success",
			fields: fields{
				repo: &fakeRepository{
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
			command, err := shomu2.NewCommand("push", tt.fields.repo)
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
