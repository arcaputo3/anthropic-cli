# Anthropic CLI

The official CLI for the [Anthropic REST API](https://docs.anthropic.com/claude/reference/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/anthropic-cli'
```

## Usage

The CLI follows a resource-based command structure:

```sh
anthropic-cli [resource] [command] [flags]
```

```sh
anthropic-cli messages create \
  --max-tokens 1024 \
  --messages.content.text x \
  --messages.content.type text \
  --messages.role user \
  --model claude-3-5-sonnet-latest
```

For details about specific commands, use the `--help` flag.
