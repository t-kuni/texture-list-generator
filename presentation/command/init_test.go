package command_test

import (
	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/di"
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(filepath.Join("..", "..", ".env.test"))
	if err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}

type TestCaseContainer struct {
	t        *testing.T
	DI       *do.Injector
	MockCtrl *gomock.Controller
}

func beforeEach(t *testing.T) *TestCaseContainer {
	return &TestCaseContainer{
		t:        t,
		DI:       di.NewContainer(),
		MockCtrl: gomock.NewController(t),
	}
}

func afterEach(cont *TestCaseContainer) {
	cont.DI.Shutdown()
	cont.MockCtrl.Finish()
}
