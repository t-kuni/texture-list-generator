package di

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/usecase"
	"github.com/t-kuni/go-cli-app-skeleton/infrastructure/system"
	"github.com/t-kuni/go-cli-app-skeleton/presentation/command"
)

func NewContainer() *do.Injector {
	injector := do.New()

	// Presentation
	do.Provide(injector, command.NewRootCommand)

	// Infrastructure
	do.Provide(injector, system.NewTimer)
	do.Provide(injector, system.NewStdio)
	do.Provide(injector, system.NewFiler)

	// UseCase
	do.Provide(injector, usecase.NewCommand1UseCase)

	// Domain

	return injector
}
