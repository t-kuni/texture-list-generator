package system

import (
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
	"time"
)

type Timer struct{}

func NewTimer(i *do.Injector) (system.ITimer, error) {
	return Timer{}, nil
}

func (t Timer) Now() time.Time {
	return time.Now()
}
