package command

import (
	_ "embed"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/t-kuni/go-cli-app-skeleton/domain/infrastructure/system"
	"html/template"
	"os"
	"path/filepath"
)

//go:embed templates/index.html
var htmlTmplByte []byte

type RootCommand struct {
	CobraCommand *cobra.Command
}

func NewRootCommand(i *do.Injector) (*RootCommand, error) {
	filer := do.MustInvoke[system.IFiler](i)

	cmd := &cobra.Command{
		Short: "RootCommand",
		Long:  `RootCommand`,
		RunE: func(cmd *cobra.Command, args []string) error {
			htmlTmpl := string(htmlTmplByte)
			t := template.Must(template.New("todos").Parse(htmlTmpl))

			paths, err := filer.FindFilePaths([]string{".png", ".jpg", ".jpeg", ".gif"})
			if err != nil {
				return err
			}

			textures := lo.Map[string, *Texture](paths, func(path string, _ int) *Texture {
				return &Texture{
					Path: path,
					Name: filepath.Base(path),
				}
			})

			err = t.Execute(os.Stdout, textures)
			if err != nil {
				return err
			}
			return nil
		},
	}

	return &RootCommand{
		CobraCommand: cmd,
	}, nil
}

type Texture struct {
	Path string
	Name string
}
