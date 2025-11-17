// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var messagesCreate = cli.Command{
	Name:  "create",
	Usage: "Send a structured list of input messages with text and/or image content, and the\nmodel will generate the next message in the conversation.",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "max-tokens",
			Usage: "The maximum number of tokens to generate before stopping.\n\nNote that our models may stop _before_ reaching this maximum. This parameter only specifies the absolute maximum number of tokens to generate.\n\nDifferent models have different maximum values for this parameter.  See [models](https://docs.claude.com/en/docs/models-overview) for details.",
		},
		&cli.StringFlag{
			Name:  "model",
			Usage: "The model that will complete your prompt.\\n\\nSee [models](https://docs.anthropic.com/en/docs/models-overview) for additional details and options.",
		},
		&cli.StringFlag{
			Name:  "service-tier",
			Usage: "Determines whether to use priority capacity (if available) or standard capacity for this request.\n\nAnthropic offers different levels of service for your API requests. See [service-tiers](https://docs.claude.com/en/api/service-tiers) for details.",
		},
		&cli.StringSliceFlag{
			Name:  "stop-sequence",
			Usage: "Custom text sequences that will cause the model to stop generating.\n\nOur models will normally stop when they have naturally completed their turn, which will result in a response `stop_reason` of `\"end_turn\"`.\n\nIf you want the model to stop generating when it encounters custom strings of text, you can use the `stop_sequences` parameter. If the model encounters one of the custom sequences, the response `stop_reason` value will be `\"stop_sequence\"` and the response `stop_sequence` value will contain the matched stop sequence.",
		},
		&cli.BoolFlag{
			Name:  "stream",
			Usage: "Whether to incrementally stream the response using server-sent events.\n\nSee [streaming](https://docs.claude.com/en/api/messages-streaming) for details.",
		},
		&cli.Float64Flag{
			Name:  "temperature",
			Usage: "Amount of randomness injected into the response.\n\nDefaults to `1.0`. Ranges from `0.0` to `1.0`. Use `temperature` closer to `0.0` for analytical / multiple choice, and closer to `1.0` for creative and generative tasks.\n\nNote that even with `temperature` of `0.0`, the results will not be fully deterministic.",
		},
		&cli.Int64Flag{
			Name:  "top-k",
			Usage: "Only sample from the top K options for each subsequent token.\n\nUsed to remove \"long tail\" low probability responses. [Learn more technical details here](https://towardsdatascience.com/how-to-sample-from-language-models-682bceb97277).\n\nRecommended for advanced use cases only. You usually only need to use `temperature`.",
		},
		&cli.Float64Flag{
			Name:  "top-p",
			Usage: "Use nucleus sampling.\n\nIn nucleus sampling, we compute the cumulative distribution over all the options for each subsequent token in decreasing probability order and cut it off once it reaches a particular probability specified by `top_p`. You should either alter `temperature` or `top_p`, but not both.\n\nRecommended for advanced use cases only. You usually only need to use `temperature`.",
		},
	},
	Action:          handleMessagesCreate,
	HideHelpCommand: true,
}

var messagesCountTokens = cli.Command{
	Name:  "count-tokens",
	Usage: "Count the number of tokens in a Message.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "model",
			Usage: "The model that will complete your prompt.\\n\\nSee [models](https://docs.anthropic.com/en/docs/models-overview) for additional details and options.",
		},
	},
	Action:          handleMessagesCountTokens,
	HideHelpCommand: true,
}

func handleMessagesCreate(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.MessageNewParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"max-tokens":     "max_tokens",
		"model":          "model",
		"service-tier":   "service_tier",
		"stop-sequences": "stop_sequences",
		"stream":         "stream",
		"temperature":    "temperature",
		"top-k":          "top_k",
		"top-p":          "top_p",
	}, &params); err != nil {
		return err
	}
	stream := client.Messages.NewStreaming(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}

func handleMessagesCountTokens(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.MessageCountTokensParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"model": "model",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Messages.CountTokens(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("messages count-tokens", json, format, transform)
}
