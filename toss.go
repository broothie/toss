package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/bobg/errors"
	"github.com/broothie/qst"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

type Toss struct {
	Requests  map[string]Request
	Responses map[string]Response
}

func New() *Toss {
	return &Toss{
		Requests:  make(map[string]Request),
		Responses: make(map[string]Response),
	}
}

func (t *Toss) RunFile(ctx context.Context, fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.Wrap(err, "opening toss file")
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	tossFile, err := decodeFile(file)
	if err != nil {
		return errors.Wrap(err, "decoding toss file")
	}

	for _, request := range tossFile.Requests {
		if err := t.RunRequest(ctx, request); err != nil {
			return errors.Wrapf(err, "running request %q", request.Name)
		}
	}

	return nil
}

func (t *Toss) RunRequest(ctx context.Context, request Request) error {
	templateCtx := &templateContext{
		Env: lo.SliceToMap(os.Environ(), func(env string) (string, string) {
			split := strings.SplitN(env, "=", 2)
			return split[0], split[1]
		}),
		Requests:  t.Requests,
		Responses: t.Responses,
	}

	path, err := templateCtx.execute("path", request.Path)
	if err != nil {
		return errors.Wrap(err, "evaluating path")
	}

	var queries qst.Pipeline
	for keyTemplate, valueTemplate := range request.Query {
		key, err := templateCtx.execute("query key", keyTemplate)
		if err != nil {
			return errors.Wrap(err, "evaluating query key")
		}

		value, err := templateCtx.execute("query value", valueTemplate)
		if err != nil {
			return errors.Wrap(err, "evaluating query value")
		}

		queries = append(queries, qst.Query(key, value))
	}

	var headers qst.Pipeline
	for keyTemplate, valueTemplate := range request.Headers {
		key, err := templateCtx.execute("header key", keyTemplate)
		if err != nil {
			return errors.Wrap(err, "evaluating header key")
		}

		value, err := templateCtx.execute("header value", valueTemplate)
		if err != nil {
			return errors.Wrap(err, "evaluating header value")
		}

		headers = append(headers, qst.Header(key, value))
	}

	req, err := qst.New(strings.ToUpper(request.Method), path,
		qst.Context(ctx),
		qst.Scheme(request.Scheme),
		qst.Host(request.Host),
		queries,
		headers,
	)
	if err != nil {
		return errors.Wrap(err, "building request")
	}

	t.Requests[request.Name] = request

	start := time.Now()
	response, err := http.DefaultClient.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		return errors.Wrap(err, "sending request")
	}

	t.Responses[request.Name] = Response{
		StatusCode: response.StatusCode,
		Headers:    lo.MapValues(response.Header, func(value []string, key string) string { return strings.Join(value, "; ") }),
	}

	fmt.Printf("%s %s %s?%s | %v %s\n", start.Format(time.RFC3339), req.Method, req.URL.Path, req.URL.RawQuery, elapsed, response.Status)
	return nil
}

func decodeFile(file *os.File) (tossFile File, err error) {
	switch extension := filepath.Ext(file.Name()); extension {
	case ".json":
		err = json.NewDecoder(file).Decode(&tossFile)

	case ".yaml", ".yml":
		err = yaml.NewDecoder(file).Decode(&tossFile)

	case ".toml":
		_, err = toml.NewDecoder(file).Decode(&tossFile)

	default:
		err = fmt.Errorf("invalid file type %q", extension)
	}

	return
}
