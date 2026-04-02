package main

import "fmt"

func main() {
	handlers := initHandlers()
	for _, h := range handlers {
		fmt.Println(h.Name)
	}
}

// Event is a constraint for event types.
type Event interface {
	EventName() string
}

// Handler is a generic event handler interface.
type Handler[T Event] interface {
	Handle(e T) string
}

// Wrapped is the common output type (like *event.AsynqHandler).
type Wrapped struct {
	Name string
}

// WrapHandler is a generic function like event.WrapHandler[T].
func WrapHandler[T Event](h Handler[T]) *Wrapped {
	e := h.Handle(*new(T))
	return &Wrapped{Name: e}
}

// -- Concrete event types --

type FooEvent struct{}

func (FooEvent) EventName() string { return "foo" }

type BarEvent struct{}

func (BarEvent) EventName() string { return "bar" }

// -- Concrete handlers --

type FooHandler struct{}

func (FooHandler) Handle(_ FooEvent) string { return "handled-foo" }

func NewFooHandler() *FooHandler { return &FooHandler{} }

type BarHandler struct{}

func (BarHandler) Handle(_ BarEvent) string { return "handled-bar" }

func NewBarHandler() *BarHandler { return &BarHandler{} }

// Handlers is the slice type.
type Handlers []*Wrapped
