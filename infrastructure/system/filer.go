package system

import (
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
	"os"
	"path/filepath"
)

type Filer struct {
}

func NewFiler(i *do.Injector) (system.IFiler, error) {
	return &Filer{}, nil
}

func (f Filer) FindFilePaths(includeExts []string) ([]string, error) {
	filePaths := []string{}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !lo.Contains(includeExts, filepath.Ext(path)) {
			return nil
		}

		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return filePaths, nil
}
