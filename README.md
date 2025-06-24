# toss

A command-line tool for sending HTTP requests from configuration files with template support.

## Overview

`toss` allows you to define HTTP requests in configuration files (JSON, YAML, or TOML) and execute them with built-in templating capabilities. It's perfect for API testing, automation scripts, and scenarios where you need to chain multiple HTTP requests together.

## Features

- **Multiple file formats**: Support for JSON, YAML, and TOML configuration files
- **Template support**: Use Go templates with environment variables and previous request/response data
- **Request chaining**: Access previous requests and responses in subsequent requests
- **Timing and logging**: Built-in request timing and response status logging
- **Environment integration**: Access environment variables in templates
- **Flexible configuration**: Support for headers, query parameters, and request bodies

## Installation

### From Source

```bash
git clone https://github.com/broothie/toss.git
cd toss
go build
```

### Using Go Install

```bash
go install github.com/broothie/toss@latest
```

## Usage

```bash
toss <config-file>
```

### Examples

```bash
# Run requests from a YAML file
toss requests.yaml

# Run requests from a JSON file  
toss api-tests.json

# Run requests from a TOML file
toss config.toml
```

## Configuration File Format

### Basic Structure

```yaml
requests:
  - name: "get-users"
    method: "GET"
    scheme: "https"
    host: "jsonplaceholder.typicode.com"
    path: "/users"
    headers:
      Accept: "application/json"
    
  - name: "get-user-posts"
    method: "GET" 
    scheme: "https"
    host: "jsonplaceholder.typicode.com"
    path: "/users/1/posts"
    query:
      _limit: "5"
```

### Request Fields

- `name`: Unique identifier for the request (used in templates)
- `method`: HTTP method (GET, POST, PUT, DELETE, etc.)
- `scheme`: URL scheme (http, https)
- `host`: Target hostname
- `path`: URL path (supports templating)
- `query`: Query parameters as key-value pairs (supports templating)
- `headers`: HTTP headers as key-value pairs (supports templating)
- `body`: Request body configuration

### Body Configuration

```yaml
requests:
  - name: "create-post"
    method: "POST"
    scheme: "https"
    host: "jsonplaceholder.typicode.com"
    path: "/posts"
    headers:
      Content-Type: "application/json"
    body:
      type: "json"
      value: |
        {
          "title": "My Post",
          "body": "Post content",
          "userId": 1
        }
```

## Template System

`toss` uses Go templates for dynamic values. You can access:

- **Environment variables**: `{{.Env.VARIABLE_NAME}}`
- **Previous requests**: `{{.Requests.request_name}}`
- **Previous responses**: `{{.Responses.request_name}}`

### Environment Variables Example

```yaml
requests:
  - name: "api-call"
    method: "GET"
    scheme: "https"
    host: "{{.Env.API_HOST}}"
    path: "/api/v1/data"
    headers:
      Authorization: "Bearer {{.Env.API_TOKEN}}"
```

```bash
export API_HOST=api.example.com
export API_TOKEN=your-token-here
toss config.yaml
```

### Request Chaining Example

```yaml
requests:
  - name: "login"
    method: "POST"
    scheme: "https"
    host: "api.example.com"
    path: "/auth/login"
    body:
      type: "json"
      value: |
        {
          "username": "{{.Env.USERNAME}}",
          "password": "{{.Env.PASSWORD}}"
        }
  
  - name: "get-profile"
    method: "GET"
    scheme: "https"
    host: "api.example.com"
    path: "/user/profile"
    headers:
      Authorization: "Bearer {{.Responses.login.Header.Get \"Authorization\"}}"
```

## File Format Examples

### JSON Format

```json
{
  "requests": [
    {
      "name": "example",
      "method": "GET",
      "scheme": "https",
      "host": "httpbin.org",
      "path": "/get",
      "query": {
        "param1": "value1"
      },
      "headers": {
        "User-Agent": "toss/1.0"
      }
    }
  ]
}
```

### TOML Format

```toml
[[requests]]
name = "example"
method = "GET"
scheme = "https"
host = "httpbin.org"
path = "/get"

[requests.query]
param1 = "value1"

[requests.headers]
User-Agent = "toss/1.0"
```

## Output Format

`toss` outputs timing and status information for each request:

```
2023-12-01T10:30:00Z GET /users? | 145ms 200 OK
2023-12-01T10:30:01Z GET /users/1/posts?_limit=5 | 89ms 200 OK
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

### Development

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests: `go test ./...`
5. Commit your changes (`git commit -m 'Add some amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

[License information to be added]

## Author  

Created by [broothie](https://github.com/broothie)