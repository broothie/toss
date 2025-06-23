package main

type File struct {
	Requests []Request `yaml:"requests"`
}

type Request struct {
	Name    string            `yaml:"name"`
	Method  string            `yaml:"method"`
	Scheme  string            `yaml:"scheme"`
	Host    string            `yaml:"host"`
	Path    string            `yaml:"path"`
	Query   map[string]string `yaml:"query"`
	Headers map[string]string `yaml:"headers"`
	Body    Body              `yaml:"body"`
}

type Body struct {
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}
