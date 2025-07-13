package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/broothie/cli"
	"gopkg.in/yaml.v3"
)

const (
	argInitName       = "name"
	flagInitDirectory = "directory"
	flagInitFileType  = "file-type"
)

func initHandler(ctx context.Context) error {
	name, err := cli.ArgValue[string](ctx, argInitName)
	if err != nil {
		return err
	}

	directory, err := cli.FlagValue[string](ctx, flagInitDirectory)
	if err != nil {
		return err
	}

	fileType, err := cli.FlagValue[string](ctx, flagInitFileType)
	if err != nil {
		return err
	}

	var fileName string
	if name != "" {
		fileName = fmt.Sprintf("%s.toss.%s", name, fileType)
	} else {
		fileName = fmt.Sprintf("toss.%s", fileType)
	}

	filePath := filepath.Join(directory, fileName)

	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file already exists: %s", filePath)
	}

	content, err := generateInitialContent(fileType)
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	fmt.Printf("Created toss file: %s\n", filePath)
	return nil
}

func generateInitialContent(fileType string) ([]byte, error) {
	file := File{
		Requests: []Request{
			{
				Name:   "example",
				Method: "GET",
				Scheme: "https",
				Host:   "httpbin.org",
				Path:   "/get",
			},
		},
	}

	switch fileType {
	case "json":
		return json.MarshalIndent(file, "", "  ")
	case "toml":
		return toml.Marshal(file)
	case "yaml", "yml":
		return yaml.Marshal(file)
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}
}
