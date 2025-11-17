# Anthropic CLI

The official CLI for the [Anthropic REST API](https://docs.anthropic.com/claude/reference/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/anthropic-cli/cmd/anthropic-cli@latest'
```

### Running Locally

```sh
go run cmd/anthropic-cli/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
anthropic-cli [resource] [command] [flags]
```

```sh
anthropic-cli messages create \
  --max-tokens 1024 \
  --model claude-sonnet-4-5-20250929 \
<<JSON
{
  "messages": [
    {
      "content": [
        {
          "text": "x",
          "type": "text"
        }
      ],
      "role": "user"
    }
  ]
}
JSON
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
