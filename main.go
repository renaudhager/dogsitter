package main

import (
	"dogsitter/commands"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var version string

// main function
func main() {
	app := cli.NewApp()
	app.Name = "dogsitter"
	app.HelpName = "CLI tool to import and export Datadog dashboard."
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key",
			Usage:  "Datadog API key",
			EnvVar: "DATADOG_API_KEY",
		},
		cli.StringFlag{
			Name:   "app-key, application-key",
			Usage:  "Datadog Application key",
			EnvVar: "DATADOG_APP_KEY",
		},
		cli.StringFlag{
			Name:   "l, log-level",
			Usage:  "Setting log level",
			EnvVar: "DS_LOGLEVEL",
			Value:  "INFO",
		},
		cli.StringFlag{
			Name:   "dh, datadog-host",
			Usage:  "Datadog endpoint",
			EnvVar: "DD_HOST",
			Value:  "https://app.datadoghq.eu",
		},
	}

	app.Commands = []cli.Command{
		commands.PullCmd,
		commands.PushCmd,
		commands.ListCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
