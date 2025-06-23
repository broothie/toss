package main

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/bobg/errors"
	"github.com/samber/lo"
)

type templateContext struct {
	Env       map[string]string
	Requests  map[string]*http.Request
	Responses map[string]*http.Response
}

func newTemplateContext(environ []string) *templateContext {
	return &templateContext{
		Env: lo.SliceToMap(environ, func(env string) (string, string) {
			split := strings.SplitN(env, "=", 2)
			return split[0], split[1]
		}),
		Requests:  make(map[string]*http.Request),
		Responses: make(map[string]*http.Response),
	}
}

func (c templateContext) execute(name string, input string) (string, error) {
	tmpl, err := template.New(name).Parse(input)
	if err != nil {
		return "", errors.Wrap(err, "parsing template")
	}

	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, c); err != nil {
		return "", errors.Wrap(err, "applying template")
	}

	return buffer.String(), nil
}
