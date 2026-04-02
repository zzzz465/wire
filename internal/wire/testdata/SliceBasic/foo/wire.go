//go:build wireinject
// +build wireinject

package main

import "github.com/goforj/wire"

func initMessages() Messages {
	wire.Build(
		wire.Slice(new(Messages),
			provideHello,
			provideWorld,
		),
	)
	return nil
}
