package system

import (
	"fmt"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
)

type Stdio struct {
}

func NewStdio(i *do.Injector) (system.IStdio, error) {
	return &Stdio{}, nil
}

func (s *Stdio) Printf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}
