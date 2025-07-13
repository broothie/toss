package main

import (
	"context"

	"github.com/broothie/cli"
)

const argFile = "file"

func tossHandler(ctx context.Context) error {
	fileName, err := cli.ArgValue[string](ctx, argFile)
	if err != nil {
		return err
	}

	toss := New()

	if err := toss.RunFile(ctx, fileName); err != nil {
		return err
	}

	return nil
}
