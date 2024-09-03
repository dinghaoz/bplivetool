package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Starting app")
	app := &cli.App{
		Name:        "bplivetool",
		HelpName:    "bplivetool",
		Description: "BytePlus Live Command Live Tool",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "access_key", Aliases: []string{"a"}, Required: true},
			&cli.StringFlag{Name: "secret_key", Aliases: []string{"s"}, Required: true},
		},
		Commands: []*cli.Command{ListCmd()},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start app")
	}
}
