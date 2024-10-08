package main

import (
	"encoding/json"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/livesaas"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func listAction(ctx *cli.Context) error {
	livesaas.DefaultInstance.Client.SetAccessKey(ctx.String("access_key"))
	livesaas.DefaultInstance.Client.SetSecretKey(ctx.String("secret_key"))

	bodyMap := map[string]interface{}{
		"PageNo":        ctx.Int("PageNo"),
		"PageItemCount": ctx.Int("PageItemCount"),
		"Name":          ctx.String("Name"),
		"Status":        ctx.Int("Status"),
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return err
	}

	log.Debug().Str("body", string(body)).Msg("request")

	resp, statusCode, err := livesaas.DefaultInstance.ListActivityAPI(nil, string(body))
	if err != nil {
		return err
	}
	log.Info().
		Int("status", statusCode).
		Msg("ListActivityAPI")

	log.Info().
		Interface("resp", resp.Result.Activities).
		Msg("ListActivityAPI")
	return nil
}

func ListCmd() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list usage",
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "PageNo", Value: 0},
			&cli.IntFlag{Name: "PageItemCount", Value: 10},
			&cli.StringFlag{Name: "Name"},
			&cli.IntFlag{Name: "Status", Value: 0},
		},
		Action: listAction,
	}
}
