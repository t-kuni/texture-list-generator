package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-skeleton/di"
	"github.com/t-kuni/go-cli-app-skeleton/presentation/command"
	"os"
)

func main() {
	ctx := context.Background()

	godotenv.Load()

	container := di.NewContainer()
	defer container.Shutdown()

	cmd := do.MustInvoke[*command.RootCommand](container)
	err := cmd.CobraCommand.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}
