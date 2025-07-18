package main

import (
	"context"

	"github.com/broothie/cli"
)

const argFile = "file"

func runHandler(ctx context.Context) error {
	fileName, err := cli.ArgValue[string](ctx, argFile)
	if err != nil {
		return err
	}

	if err := New().RunFile(ctx, fileName); err != nil {
		return err
	}

	return nil
}
