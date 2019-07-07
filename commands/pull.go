package commands

import (
	"bytes"
	"encoding/json"
	"errors"
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

	if err != nil {
		log.Error("Error when querying Datadog API.")
		return err
	}

	if statusCode != 200 {
		log.Error("Return is not 200: ", statusCode)
		return errors.New("Return code is not 200")
	}

	content := beautify(payload)

	if c.String("o") == "stdout" {
		fmt.Print(string(content))
		fmt.Print("\n")
	} else {
		err = dumpDashboard(content, c.String("o"))
		if err != nil {
			return errors.New("Cannot dump the dashboard")
		}
	}

	return nil
}

// getDashboard get dashboard JSON payload from Datadog
func getDashboard(ddEndpoint string, dashboardID string, apiKey string, appKey string) (string, int, error) {

	var (
		body       []byte
		query      string
		statusCode int
	)

	log.Info("Pulling dashboard ", dashboardID)

	query = ddEndpoint + "/api/v1/dashboard/" + dashboardID + "?api_key=" + apiKey + "&application_key=" + appKey

	resp, err := http.Get(query)

	if err != nil {
		log.Error("Error connection to ", query)
		return "", 500, err
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode

	if statusCode != 200 {
		log.Error("Returned code is not 200, it's ", resp.StatusCode)
		return string(body), statusCode, nil
	}

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Unable to read body of the response")
		return "", 500, err
	}

	// Call to a function to strip a field that breaks Datadog API when the JSON is imported
	strippedPayload, _ := stripBadField(body, "author_name")

	return string(strippedPayload), statusCode, err
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

func dumpDashboard(content []byte, filepath string) error {

	err := ioutil.WriteFile(filepath, content, 0600)
	if err != nil {
		log.Error("Error when writing to ", filepath)
		return err
	}

	log.Info("Dashboard dumped into ", filepath)

	return err
}

// stripBadField function is required because Datadog is broken.
// We need to remove field `author_name` from the payload.
func stripBadField(payload []byte, pattern string) ([]byte, error) {

	var strippedPayload []byte

	m := make(map[string]interface{})
	n := make(map[string]interface{})

	err := json.Unmarshal(payload, &m)

	if err != nil {
		log.Error("Error when unmarshalling, ", err)
		return strippedPayload, err
	}

	for k, v := range m {
		// Removing key pattern from the payload
		if k != pattern {
			n[k] = v
		}
	}

	strippedPayload, err = json.Marshal(n)
	if err != nil {
		log.Error("Error when marshalling, ", err)
		return strippedPayload, err
	}

	return strippedPayload, nil
}
