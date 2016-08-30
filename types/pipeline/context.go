package pipeline

type Context interface {
}

type context struct {
}

func NewContext() Context {
	return &context{}
}
