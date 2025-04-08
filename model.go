// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func createModelsRetrieveSubcommand() Subcommand {
	var modelID *string = nil
	query := []byte("{}")
	header := []byte("{}")
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
			res, err := client.Models.Get(
				context.TODO(),
				*modelID,
				option.WithMiddleware(func(r *http.Request, mn option.MiddlewareNext) (*http.Response, error) {
					q := r.URL.Query()
					for key, values := range serializeQuery(query) {
						for _, value := range values {
							q.Add(key, value)
						}
					}
					r.URL.RawQuery = q.Encode()

					for key, values := range serializeHeader(header) {
						for _, value := range values {
							r.Header.Add(key, value)
						}
					}

					return mn(r)
				}),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createModelsListSubcommand() Subcommand {
	query := []byte("{}")
	header := []byte("{}")
	var flagSet = flag.NewFlagSet("models.list", flag.ExitOnError)

	flagSet.Func(
		"after-id",
		"",
		func(string string) error {
			var jsonErr error
			query, jsonErr = jsonSet(query, "after_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"before-id",
		"",
		func(string string) error {
			var jsonErr error
			query, jsonErr = jsonSet(query, "before_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"limit",
		"",
		func(string string) error {
			int, err := parseInt(string)
			if err != nil {
				return err
			}
			var jsonErr error
			query, jsonErr = jsonSet(query, "limit", int)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Models.List(
				context.TODO(),
				anthropic.ModelListParams{},
				option.WithMiddleware(func(r *http.Request, mn option.MiddlewareNext) (*http.Response, error) {
					q := r.URL.Query()
					for key, values := range serializeQuery(query) {
						for _, value := range values {
							q.Add(key, value)
						}
					}
					r.URL.RawQuery = q.Encode()

					for key, values := range serializeHeader(header) {
						for _, value := range values {
							r.Header.Add(key, value)
						}
					}

					return mn(r)
				}),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}
