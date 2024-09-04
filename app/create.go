package main

import (
	"encoding/json"
	"github.com/byteplus-sdk/byteplus-sdk-golang/service/livesaas"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func createAction(ctx *cli.Context) error {
	livesaas.DefaultInstance.Client.SetAccessKey(ctx.String("access_key"))
	livesaas.DefaultInstance.Client.SetSecretKey(ctx.String("secret_key"))

	bodyMap := map[string]interface{}{
		"CoverImage": ctx.String("CoverImage"),
		"TemplateId": ctx.Int64("TemplateId"),
		"Name":       ctx.String("Name"),
		"LiveMode":   ctx.Int("LiveMode"),
	}
	body, err := json.Marshal(bodyMap)
	if err != nil {
		return err
	}

	log.Debug().Str("body", string(body)).Msg("request")

	resp, statusCode, err := livesaas.DefaultInstance.CreateActivityAPIV2(nil, string(body))
	if err != nil {
		return err
	}
	log.Info().
		Int("status", statusCode).
		Msg("CreateActivityAPIV2")

	log.Info().
		Interface("resp", resp.Result).
		Msg("CreateActivityAPIV2")
	return nil
}

func CreateCmd() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "create usage",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "CoverImage"},
			&cli.Int64Flag{Name: "TemplateId"},
			&cli.StringFlag{Name: "Name"},
			&cli.IntFlag{Name: "LiveMode", Value: 0},
		},
		Action: createAction,
	}
}
