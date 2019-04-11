package utils

import (

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
}

// PushCmd pull to download dashboard configuration from Datadog.
var PushCmd = cli.Command{
	Name:   "push",
	Usage:  "Import dashboard configuration to Datadog.",
	Action: push,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "f, file",
			Usage:  "Dashboard file configuration.",
			EnvVar: "DS_IMPORT_FILE",
		},
	},
}

func push(c *cli.Context) (err error) {

	return err
}
