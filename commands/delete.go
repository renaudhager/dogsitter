package commands

import (
	"bytes"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Delete interface
type Delete interface {
	deleteDashboard(c *cli.Context) error
}

// DeleteAction struct
type DeleteAction struct{}

// NewDeleteAction constructor for DeleteAction
func NewDeleteAction() Delete {
	return &DeleteAction{}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// DeleteCmd pull to download dashboard configuration from Datadog.
var DeleteCmd = cli.Command{
	Name:   "delete",
	Usage:  "Delete dashboard from Datadog.",
	Action: actionDelete,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "id",
			Usage: "Dashboard id.",
		},
	},
}

// actionDelete placeholder function
func actionDelete(c *cli.Context) (err error) {
	return NewDeleteAction().deleteDashboard(c)
}

func (da *DeleteAction) deleteDashboard(c *cli.Context) error {

	ddEndpoint := c.GlobalString("dh")
	dasboardID := c.String("id")
	apiKey := c.GlobalString("api-key")
	appKey := c.GlobalString("app-key")

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

	log.Info("Successfully deleted dashboard ", dasboardID)
	return nil
}
