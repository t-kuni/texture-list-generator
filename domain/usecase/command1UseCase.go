package usecase

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
)

type Command1UseCase struct {
	stdio system.IStdio
	filer system.IFiler
}

func NewCommand1UseCase(i *do.Injector) (*Command1UseCase, error) {
	return &Command1UseCase{
		stdio: do.MustInvoke[system.IStdio](i),
		filer: do.MustInvoke[system.IFiler](i),
	}, nil
}

func (u Command1UseCase) Execute() error {
	return nil
}
