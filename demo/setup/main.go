package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	container := client.Container().
		From("golang:1.19").
		WithExec([]string{"echo", "hello"})

	if _, err := container.ExitCode(ctx); err != nil {
		panic(err)
	}
}
