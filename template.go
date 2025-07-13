package main

import (
	"bytes"
	"html/template"

	"github.com/bobg/errors"
)

type templateContext struct {
	Env       map[string]string
	Requests  map[string]Request
	Responses map[string]Response
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
