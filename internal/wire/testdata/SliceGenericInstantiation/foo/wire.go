//go:build wireinject
// +build wireinject

package main

import "github.com/goforj/wire"

// Reproduce GSM pattern: constructors in a separate Set, Slice in another Set that imports it.
var ConstructorSet = wire.NewSet(
	NewFooHandler,
	NewBarHandler,
)

var SliceSet = wire.NewSet(
	ConstructorSet,
	wire.Slice(new(Handlers),
		WrapHandler[FooEvent],
		WrapHandler[BarEvent],
	),
)

func initHandlers() Handlers {
	wire.Build(SliceSet)
	return nil
}
