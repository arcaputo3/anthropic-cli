// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/stainless-sdks/anthropic-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var betaFilesList = cli.Command{
	Name:  "list",
	Usage: "List Files",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "after-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "after_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "before-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "before_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "betas",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+beta",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaFilesList,
	HideHelpCommand: true,
}

var betaFilesDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete File",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "file-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "betas",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+beta",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaFilesDelete,
	HideHelpCommand: true,
}

var betaFilesDownload = cli.Command{
	Name:  "download",
	Usage: "Download File",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "file-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "betas",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+beta",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaFilesDownload,
	HideHelpCommand: true,
}

var betaFilesRetrieveMetadata = cli.Command{
	Name:  "retrieve-metadata",
	Usage: "Get File Metadata",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "file-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "betas",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+beta",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaFilesRetrieveMetadata,
	HideHelpCommand: true,
}

var betaFilesUpload = cli.Command{
	Name:  "upload",
	Usage: "Upload File",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "file",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "file",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "betas",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+beta",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaFilesUpload,
	HideHelpCommand: true,
}

func handleBetaFilesList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := anthropic.BetaFileListParams{}
	var res []byte
	_, err := cc.client.Beta.Files.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("beta:files list", string(res), format)
}

func handleBetaFilesDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := anthropic.BetaFileDeleteParams{}
	var res []byte
	_, err := cc.client.Beta.Files.Delete(
		context.TODO(),
		cmd.Value("file-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("beta:files delete", string(res), format)
}

func handleBetaFilesDownload(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := anthropic.BetaFileDownloadParams{}
	var res []byte
	_, err := cc.client.Beta.Files.Download(
		context.TODO(),
		cmd.Value("file-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("beta:files download", string(res), format)
}

func handleBetaFilesRetrieveMetadata(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := anthropic.BetaFileGetMetadataParams{}
	var res []byte
	_, err := cc.client.Beta.Files.GetMetadata(
		context.TODO(),
		cmd.Value("file-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("beta:files retrieve-metadata", string(res), format)
}

func handleBetaFilesUpload(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := anthropic.BetaFileUploadParams{}
	var res []byte
	_, err := cc.client.Beta.Files.Upload(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("beta:files upload", string(res), format)
}
