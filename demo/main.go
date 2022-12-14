package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// クライアントを初期化して Dagger Engine に接続する
	// dagger.WithLogOutput でログの出力先を指定できる
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Docker イメージを取得する
	container := client.Container().From("golang:1.19")

	// カレントディレクトリをコンテナにマウントする
	src := client.Host().Directory(".")
	container = container.
		WithMountedDirectory("/src", src).
		WithWorkdir("/src")

	// 実行するコマンドを設定する
	container = container.
		WithExec([]string{"go", "test", "-v", "./..."}).
		WithExec([]string{"go", "build"})

	// パイプラインを実行する
	if _, err := container.ExitCode(ctx); err != nil {
		panic(err)
	}
}
