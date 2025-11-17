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

var messagesBatchesCreate = cli.Command{
	Name:            "create",
	Usage:           "Send a batch of Message creation requests.",
	Flags:           []cli.Flag{},
	Action:          handleMessagesBatchesCreate,
	HideHelpCommand: true,
}

var messagesBatchesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "This endpoint is idempotent and can be used to poll for Message Batch\ncompletion. To access the results of a Message Batch, make a request to the\n`results_url` field in the response.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "message-batch-id",
			Usage: "ID of the Message Batch.",
		},
	},
	Action:          handleMessagesBatchesRetrieve,
	HideHelpCommand: true,
}

var messagesBatchesList = cli.Command{
	Name:  "list",
	Usage: "List all Message Batches within a Workspace. Most recently created batches are\nreturned first.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "after-id",
			Usage: "ID of the object to use as a cursor for pagination. When provided, returns the page of results immediately after this object.",
		},
		&cli.StringFlag{
			Name:  "before-id",
			Usage: "ID of the object to use as a cursor for pagination. When provided, returns the page of results immediately before this object.",
		},
		&cli.Int64Flag{
			Name:  "limit",
			Usage: "Number of items to return per page.\n\nDefaults to `20`. Ranges from `1` to `1000`.",
			Value: 20,
		},
	},
	Action:          handleMessagesBatchesList,
	HideHelpCommand: true,
}

var messagesBatchesDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a Message Batch.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "message-batch-id",
			Usage: "ID of the Message Batch.",
		},
	},
	Action:          handleMessagesBatchesDelete,
	HideHelpCommand: true,
}

var messagesBatchesCancel = cli.Command{
	Name:  "cancel",
	Usage: "Batches may be canceled any time before processing ends. Once cancellation is\ninitiated, the batch enters a `canceling` state, at which time the system may\ncomplete any in-progress, non-interruptible requests before finalizing\ncancellation.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "message-batch-id",
			Usage: "ID of the Message Batch.",
		},
	},
	Action:          handleMessagesBatchesCancel,
	HideHelpCommand: true,
}

var messagesBatchesResults = cli.Command{
	Name:  "results",
	Usage: "Streams the results of a Message Batch as a `.jsonl` file.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "message-batch-id",
			Usage: "ID of the Message Batch.",
		},
	},
	Action:          handleMessagesBatchesResults,
	HideHelpCommand: true,
}

func handleMessagesBatchesCreate(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.MessageBatchNewParams{}
	if err := unmarshalStdinWithFlags(cmd, nil, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.Messages.Batches.New(
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
	return ShowJSON("messages:batches create", json, format, transform)
}

func handleMessagesBatchesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-batch-id") && len(unusedArgs) > 0 {
		cmd.Set("message-batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Messages.Batches.Get(
		ctx,
		cmd.Value("message-batch-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("messages:batches retrieve", json, format, transform)
}

func handleMessagesBatchesList(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.MessageBatchListParams{
		AfterID:  anthropic.String(cmd.Value("after-id").(string)),
		BeforeID: anthropic.String(cmd.Value("before-id").(string)),
	}
	if cmd.IsSet("limit") {
		params.Limit = anthropic.Opt(cmd.Value("limit").(int64))
	}
	var res []byte
	_, err := client.Messages.Batches.List(
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
	return ShowJSON("messages:batches list", json, format, transform)
}

func handleMessagesBatchesDelete(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-batch-id") && len(unusedArgs) > 0 {
		cmd.Set("message-batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Messages.Batches.Delete(
		ctx,
		cmd.Value("message-batch-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("messages:batches delete", json, format, transform)
}

func handleMessagesBatchesCancel(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-batch-id") && len(unusedArgs) > 0 {
		cmd.Set("message-batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.Messages.Batches.Cancel(
		ctx,
		cmd.Value("message-batch-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("messages:batches cancel", json, format, transform)
}

func handleMessagesBatchesResults(ctx context.Context, cmd *cli.Command) error {
	client := anthropic.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("message-batch-id") && len(unusedArgs) > 0 {
		cmd.Set("message-batch-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	stream := client.Messages.Batches.ResultsStreaming(
		ctx,
		cmd.Value("message-batch-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}
