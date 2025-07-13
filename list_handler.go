package main

import (
	"context"
	"fmt"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/bobg/errors"
)

func listHandler(context.Context) error {
	patterns := []string{
		"**/toss.{yml,yaml,json,toml}",
		"**/*.toss.{yml,yaml,json,toml}",
	}

	for _, pattern := range patterns {
		matches, err := doublestar.FilepathGlob(pattern)
		if err != nil {
			return errors.Wrap(err, "globbing for toss files")
		}
		for _, match := range matches {
			fmt.Println(match)
		}
	}

	return nil
}
