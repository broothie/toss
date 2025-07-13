package main

type File struct {
	Requests []Request `yaml:"requests" json:"requests" toml:"requests"`
}

type Request struct {
	Name    string            `yaml:"name" json:"name" toml:"name"`
	Method  string            `yaml:"method" json:"method" toml:"method"`
	Scheme  string            `yaml:"scheme" json:"scheme" toml:"scheme"`
	Host    string            `yaml:"host" json:"host" toml:"host"`
	Path    string            `yaml:"path" json:"path" toml:"path"`
	Query   map[string]string `yaml:"query" json:"query" toml:"query"`
	Headers map[string]string `yaml:"headers" json:"headers" toml:"headers"`
	Body    Body              `yaml:"body" json:"body" toml:"body"`
}

type Body struct {
	Type  string `yaml:"type" json:"type" toml:"type"`
	Value string `yaml:"value" json:"value" toml:"value"`
}

type Response struct {
	StatusCode int               `yaml:"status_code" json:"status_code" toml:"status_code"`
	Headers    map[string]string `yaml:"headers" json:"headers" toml:"headers"`
}
