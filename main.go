// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand")
		os.Exit(1)
	}

	subcommand := subcommands[os.Args[1]]
	if subcommand == nil {
		log.Fatalf("Unknown subcommand '%s'", os.Args[1])
	}

	subcommand.flagSet.Parse(os.Args[2:])

	var client *anthropic.Client = anthropic.NewClient()
	subcommand.handle(client)
}

func init() {
	initialJson := getStdInput()
	if initialJson == nil {
		initialJson = []byte("{}")
	}

	var completionsCreateSubcommand = createCompletionsCreateSubcommand(initialJson)
	subcommands[completionsCreateSubcommand.flagSet.Name()] = &completionsCreateSubcommand

	var messagesCreateSubcommand = createMessagesCreateSubcommand(initialJson)
	subcommands[messagesCreateSubcommand.flagSet.Name()] = &messagesCreateSubcommand

	var messagesCountTokensSubcommand = createMessagesCountTokensSubcommand(initialJson)
	subcommands[messagesCountTokensSubcommand.flagSet.Name()] = &messagesCountTokensSubcommand

	var messagesBatchesCreateSubcommand = createMessagesBatchesCreateSubcommand(initialJson)
	subcommands[messagesBatchesCreateSubcommand.flagSet.Name()] = &messagesBatchesCreateSubcommand

	var messagesBatchesRetrieveSubcommand = createMessagesBatchesRetrieveSubcommand()
	subcommands[messagesBatchesRetrieveSubcommand.flagSet.Name()] = &messagesBatchesRetrieveSubcommand

	var messagesBatchesListSubcommand = createMessagesBatchesListSubcommand()
	subcommands[messagesBatchesListSubcommand.flagSet.Name()] = &messagesBatchesListSubcommand

	var messagesBatchesDeleteSubcommand = createMessagesBatchesDeleteSubcommand(initialJson)
	subcommands[messagesBatchesDeleteSubcommand.flagSet.Name()] = &messagesBatchesDeleteSubcommand

	var messagesBatchesCancelSubcommand = createMessagesBatchesCancelSubcommand(initialJson)
	subcommands[messagesBatchesCancelSubcommand.flagSet.Name()] = &messagesBatchesCancelSubcommand

	var modelsRetrieveSubcommand = createModelsRetrieveSubcommand()
	subcommands[modelsRetrieveSubcommand.flagSet.Name()] = &modelsRetrieveSubcommand

	var modelsListSubcommand = createModelsListSubcommand()
	subcommands[modelsListSubcommand.flagSet.Name()] = &modelsListSubcommand
}

var subcommands = map[string]*Subcommand{}

func createCompletionsCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("completions.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens-to-sample",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "max_tokens_to_sample", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "model", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"prompt",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "prompt", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "metadata.user_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.#", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.-1", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "temperature", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_k", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_p", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			var err error
			json, err = jsonSet(json, "stream", true)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			stream := client.Completions.NewStreaming(
				context.TODO(),
				anthropic.CompletionNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}

func createMessagesCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "max_tokens", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"messages.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.tool_use_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "messages.#.content.#.is_error", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+document_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.role",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.role", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+message",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "model", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "metadata.user_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.#", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "stop_sequences.-1", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"system.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+system",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "system.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "temperature", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"tool-choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"tools.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.description",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.description", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+tool",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "tools.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_k", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "top_p", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			var err error
			json, err = jsonSet(json, "stream", true)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			stream := client.Messages.NewStreaming(
				context.TODO(),
				anthropic.MessageNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}

func createMessagesCountTokensSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.count_tokens", flag.ExitOnError)

	flagSet.Func(
		"messages.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.tool_use_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "messages.#.content.#.is_error", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+document_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"messages.role",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "messages.#.role", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+message",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "messages.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "model", string)
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"system.union_member_0",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"system.union_member_1.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "system.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"system.+union_member_1",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "system.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	flagSet.Func(
		"tool-choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)

	flagSet.Func(
		"tools.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"tools.description",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "tools.#.description", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+tool",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "tools.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.CountTokens(
				context.TODO(),
				anthropic.MessageCountTokensParams{},
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.create", flag.ExitOnError)

	flagSet.Func(
		"requests.custom_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.custom_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.max_tokens",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.max_tokens", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.tool_use_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.messages.#.content.#.is_error", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.data", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.media_type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.source.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.document_block_param.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.+text_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.+image_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.messages.content.+document_block_param",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.messages.role",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.#.role", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.+message",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.messages.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.model",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.model", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.metadata.user_id",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.metadata.user_id", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.stop_sequences",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.stop_sequences.#", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.+stop_sequence",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.stop_sequences.-1", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.stream",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.stream", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.system.text",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.system.#.text", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.system.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.system.#.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.system.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.system.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.+system",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.system.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.temperature", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_any.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tool_choice.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tool_choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", boolean)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tools.name",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tools.#.name", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tools.cache_control.type",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tools.#.cache_control.type", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.tools.description",
		"",
		func(string string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tools.#.description", string)
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"requests.params.+tool",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.#.params.tools.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.top_k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.top_k", integer)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.Func(
		"requests.params.top_p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			var err2 error
			json, err2 = jsonSet(json, "requests.#.params.top_p", number)
			if err2 != nil {
				return err2
			}
			return nil
		},
	)
	flagSet.BoolFunc(
		"+request",
		"",
		func(_ string) error {
			var err error
			json, err = jsonSet(json, "requests.-1", map[string]interface{}{})
			if err != nil {
				return err
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.New(
				context.TODO(),
				anthropic.MessageBatchNewParams{},
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesRetrieveSubcommand() Subcommand {
	var messageBatchID *string = nil
	var flagSet = flag.NewFlagSet("messages.batches.retrieve", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Get(context.TODO(), *messageBatchID)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesListSubcommand() Subcommand {
	var flagSet = flag.NewFlagSet("messages.batches.list", flag.ExitOnError)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.List(context.TODO(), anthropic.MessageBatchListParams{})
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesDeleteSubcommand(initialJson []byte) Subcommand {
	var messageBatchID *string = nil
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.delete", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Delete(
				context.TODO(),
				*messageBatchID,
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesCancelSubcommand(initialJson []byte) Subcommand {
	var messageBatchID *string = nil
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.cancel", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Cancel(
				context.TODO(),
				*messageBatchID,
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createModelsRetrieveSubcommand() Subcommand {
	var modelID *string = nil
	var flagSet = flag.NewFlagSet("models.retrieve", flag.ExitOnError)

	flagSet.Func(
		"model-id",
		"",
		func(string string) error {
			modelID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Models.Get(context.TODO(), *modelID)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createModelsListSubcommand() Subcommand {
	var flagSet = flag.NewFlagSet("models.list", flag.ExitOnError)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Models.List(context.TODO(), anthropic.ModelListParams{})
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

type Subcommand struct {
	flagSet *flag.FlagSet
	handle  func(*anthropic.Client)
}

func parseInt(string string) (int64, error) {
	integer, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return 0, err
	}

	return integer, nil
}

func parseFloat(string string) (float64, error) {
	number, err := strconv.ParseFloat(string, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func jsonSet(json []byte, path string, value interface{}) ([]byte, error) {
	keys := strings.Split(path, ".")
	path = ""
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if key == "#" {
			key = strconv.Itoa(len(gjson.GetBytes(json, path).Array()) - 1)
		}

		if len(path) > 0 {
			path += "."
		}
		path += key
	}
	return sjson.SetBytes(json, path, value)
}

func getStdInput() []byte {
	if !isInputPiped() {
		return nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
