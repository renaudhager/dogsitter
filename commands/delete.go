package commands

import (
	"bytes"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// DeleteCmd pull to download dashboard configuration from Datadog.
var DeleteCmd = cli.Command{
	Name:   "delete",
	Usage:  "Delete dashboard from Datadog.",
	Action: delete,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "Dashboard id.",
		},
	},
}

func delete(c *cli.Context) (err error) {

	_ = deleteDashboard(c.GlobalString("dh"), c.String("id"), c.GlobalString("api-key"), c.GlobalString("app-key"))
	return err
}

func deleteDashboard(ddEndpoint string, dasboardID string, apiKey string, appKey string) error {

	client := http.Client{}

	url := ddEndpoint + "/api/v1/dashboard/" + dasboardID + "?api_key=" + apiKey + "&application_key=" + appKey
	request, err := http.NewRequest("DELETE", url, bytes.NewBufferString(""))

	if err != nil {
		log.Error("Error creating the request to ", err)
		return err
	}

	log.Info("Deleting dashboard ", dasboardID)
	resp, err := client.Do(request)

	if err != nil {
		log.Error("Error while querying Datadog ", err)
		return err
	}

	defer resp.Body.Close()

	statusCode := resp.StatusCode

	if statusCode != 200 {
		log.Error("Returned code is not 200, it's ", statusCode)
		return errors.New("Returned code is not 200")
	}

	log.Info("Succesfully deleted dashboard ", dasboardID)
	return nil
}
