// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
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
	initialBody := getStdInput()
	if initialBody == nil {
		initialBody = []byte("{}")
	}

	var completionsCreateSubcommand = createCompletionsCreateSubcommand(initialBody)
	subcommands[completionsCreateSubcommand.flagSet.Name()] = &completionsCreateSubcommand

	var messagesCreateSubcommand = createMessagesCreateSubcommand(initialBody)
	subcommands[messagesCreateSubcommand.flagSet.Name()] = &messagesCreateSubcommand

	var messagesCountTokensSubcommand = createMessagesCountTokensSubcommand(initialBody)
	subcommands[messagesCountTokensSubcommand.flagSet.Name()] = &messagesCountTokensSubcommand

	var messagesBatchesCreateSubcommand = createMessagesBatchesCreateSubcommand(initialBody)
	subcommands[messagesBatchesCreateSubcommand.flagSet.Name()] = &messagesBatchesCreateSubcommand

	var messagesBatchesRetrieveSubcommand = createMessagesBatchesRetrieveSubcommand()
	subcommands[messagesBatchesRetrieveSubcommand.flagSet.Name()] = &messagesBatchesRetrieveSubcommand

	var messagesBatchesListSubcommand = createMessagesBatchesListSubcommand()
	subcommands[messagesBatchesListSubcommand.flagSet.Name()] = &messagesBatchesListSubcommand

	var messagesBatchesDeleteSubcommand = createMessagesBatchesDeleteSubcommand(initialBody)
	subcommands[messagesBatchesDeleteSubcommand.flagSet.Name()] = &messagesBatchesDeleteSubcommand

	var messagesBatchesCancelSubcommand = createMessagesBatchesCancelSubcommand(initialBody)
	subcommands[messagesBatchesCancelSubcommand.flagSet.Name()] = &messagesBatchesCancelSubcommand

	var modelsRetrieveSubcommand = createModelsRetrieveSubcommand()
	subcommands[modelsRetrieveSubcommand.flagSet.Name()] = &modelsRetrieveSubcommand

	var modelsListSubcommand = createModelsListSubcommand()
	subcommands[modelsListSubcommand.flagSet.Name()] = &modelsListSubcommand
}

var subcommands = map[string]*Subcommand{}

type Subcommand struct {
	flagSet *flag.FlagSet
	handle  func(*anthropic.Client)
}
