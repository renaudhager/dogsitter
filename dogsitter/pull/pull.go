package pull

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
}

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
			Usage: "output file for JSON payload.",
			Value: "stdout",
		},
	},
}

func pull(c *cli.Context) (err error) {

	payload, statusCode, err := getDashboard(c.GlobalString("dh"), c.String("id"), c.GlobalString("api-key"), c.GlobalString("app-key"))

	if err != nil  {
		log.Error("Error when querying Datadog API.")
	} else if statusCode != 200 {
		log.Error("Return is not 200: ", statusCode)
	} else {
		content := beautify(payload)

		if c.String("o") == "stdout" {
			fmt.Print(string(content))
			fmt.Print("\n")
		} else {
			dumpDashboard(content, c.String("o"))
		}

	}

	return err
}

// getDashboard get dashboard JSON payload from Datadog
func getDashboard(ddEndpoint string, dashboardID string, apiKey string, appKey string) (string, int, error) {

	var (
		body 				[]byte
		query 			string
		statusCode int
	)

	log.Info("Pulling dashboard ", dashboardID)

	query = ddEndpoint + "/api/v1/dashboard/" + dashboardID + "?api_key=" + apiKey + "&application_key=" + appKey

	resp, err := http.Get(query)
	defer resp.Body.Close()

	if err != nil {
		log.Error("Error connectiong to ", query)
	} else {
		statusCode = resp.StatusCode
		if statusCode == 200 {
			body, err = ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Error("Unable to read body of the repsonse")
			}

		} else {
			log.Error("Returned code is not 200, it's ", resp.StatusCode)
		}
	}

	return string(body), statusCode, err
}

// beautify JSON payload
func beautify(payload string) []byte {

	var (
		output     []byte
		prettyJSON bytes.Buffer
	)

	// Try to beautify JSON payload
  // if this failed dump the raw payload
	err := json.Indent(&prettyJSON, []byte(payload), "", "\t")
	if err == nil {
		output = prettyJSON.Bytes()
	} else {
		log.Warn("JSON parse error: ", err)
		output = []byte(payload)
	}

	return output
}

func dumpDashboard(content []byte, filepath string) {

	err := ioutil.WriteFile(filepath, content, 0600)
	if err != nil {
		log.Error("Error when writing to ", filepath)
	} else {
    log.Info("Dashboard dumped into ",filepath)
  }

}
