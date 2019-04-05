package pull

import (
  "io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// PullCmd pull to download dashboard configuration from Datadog.
var PullCmd = cli.Command{
	Name:   "pull",
	Usage:  "Pull dashboard configuration from Datadog API",
	Action: pull,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "id of dashboard",
		},
		cli.StringFlag{
			Name:  "o, output",
			Usage: "output file for JSON payload. If not specified title of dashboard will be used as filename(space will be replace by _).",
		},
	},
}

func pull(c *cli.Context) (err error) {

	log := confLogging(c.GlobalString("l"))
  dashboardID := c.String("id")

	log.Info("Pulling dashboard ", dashboardID)

  query := c.GlobalString("dh") + "/api/v1/dashboard/" + dashboardID + "?api_key=" + c.GlobalString("api-key") + "&application_key=" + c.GlobalString("app-key")

	resp, err := http.Get(query)
	if err != nil {
    log.Fatal("Error connectiong to ", query)
	}
	defer resp.Body.Close()

  if resp.StatusCode == 200 {
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
      log.Fatal("Unable to read body of the repsonse")
    }

    dumpDashboard(string(body), c.String("o"))

  } else {
    log.Error("Returned code is not 200, it's ", resp.StatusCode)
  }

	return nil
}

func dumpDashboard(content string, filepath string) {

  err := ioutil.WriteFile(filepath, []byte(content), 0600)
	if err != nil {
		log.Error("Error when writing to ", filepath)
	}

}

func confLogging(level string) *log.Logger {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{})
	switch level {
	case "INFO":
		logger.SetLevel(log.InfoLevel)
		logger.Info("Log level is ", level)
	case "DEBUG":
		logger.SetLevel(log.DebugLevel)
		logger.Debug("Log level is ", level)
	case "WARN":
		logger.SetLevel(log.WarnLevel)
		logger.Warn("Log level is ", level)
	case "ERROR":
		logger.SetLevel(log.ErrorLevel)
		logger.Error("Log level is ", level)
	}

	return logger
}
