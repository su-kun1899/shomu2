package shomu2

type Command interface {
	Run(options ...string)
}

type Push struct {
}

func (*Push) Run(options ...string) {
	panic("implement me")
}

func NewCommand(name string)(Command, error) {
	return &Push{}, nil
}
