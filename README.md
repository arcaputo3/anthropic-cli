# Claude Developer Platform CLI

The official CLI for the [Anthropic REST API](https://docs.anthropic.com/claude/reference/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/anthropic-cli/cmd/cdp@latest'
```

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
cdp [resource] [command] [flags]
```

```sh
cdp messages create \
  --max-tokens 1024 \
  --message '{content: [{text: x, type: text, cache_control: {type: ephemeral, ttl: 5m}, citations: [{cited_text: cited_text, document_index: 0, document_title: x, end_char_index: 0, start_char_index: 0, type: char_location}]}], role: user}' \
  --model claude-sonnet-4-5-20250929 \
  --metadata '{user_id: 13803d75-b4b5-4c3e-b2a2-6f21399b021b}' \
  --service-tier auto \
  --stop-sequence string \
  --stream false \
  --system "{text: Today's date is 2024-06-01., type: text, cache_control: {type: ephemeral, ttl: 5m}, citations: [{cited_text: cited_text, document_index: 0, document_title: x, end_char_index: 0, start_char_index: 0, type: char_location}]}" \
  --temperature 1 \
  --thinking '{budget_tokens: 1024, type: enabled}' \
  --tool-choice '{type: auto, disable_parallel_tool_use: true}' \
  --tool '{name: name, cache_control: {type: ephemeral, ttl: 5m}, description: Get the current weather in a given location, type: custom}' \
  --top-k 5 \
  --top-p 0.7
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
