package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	userAgent = "Dogsitter"
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

	content, err := loadDashboard(c.String("f"))
	if err != nil {
		log.Error("Unable to load file ", c.String("f"))
	}

	uploadDashboard(c.GlobalString("dh"), content, c.GlobalString("api-key"), c.GlobalString("app-key"))

	return err
}

// loadDashboard load a file into a string
func loadDashboard(filepath string) ([]byte, error) {

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		log.Error("Error when loading file ", filepath)
	} else {
		log.Info("Dashboard configuration loaded ", filepath)
	}

	return content, err
}

// uploadDashboard upload dashboard to Datadog
func uploadDashboard(endpoint string, content []byte, apiKey string, appKey string) (err error) {

	var datadogSite string

	url := endpoint + "/api/v1/dashboard?api_key=" + apiKey + "&application_key=" + appKey
	contentIO := bytes.NewBuffer(content)

	req, err := http.NewRequest("POST", url, contentIO)
	if err != nil {
		log.Error("Error during POST request creation: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error executing POST request:", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	statusCode := resp.StatusCode

	if statusCode != 200 {
		log.Error("Error returned is not 200:", statusCode)
		log.Errorf("Response returned by Datadog: \n %s \n", string(body))
		err = errors.New("Error during dashboard import")
		return err
	}

	// Extracting url and id of newly created dashboard
	dashboardID, dasboardURL, err := getDashboardInfo(string(body))

	if err != nil {
		log.Error("Error when parsing returned response:", err)
		return err
	}

	if strings.Contains(endpoint, "eu") {
		datadogSite = "datadoghq.eu"
	} else {
		datadogSite = "datadoghq.com"
	}

	log.Info("Dashboard successfully imported.")
	log.Infof("Dashboard id: %s", dashboardID)
	log.Infof("Dashboard url: %s", "https://"+datadogSite+dasboardURL)

	return err
}

// getDashboardInfo function thaty parse the JOSN returned by Datadog and extract
// id and url of the new dashboard
func getDashboardInfo(dashboard string) (string, string, error) {

	var (
		id  string
		url string
	)

	m := map[string]interface{}{}

	err := json.Unmarshal([]byte(dashboard), &m)
	if err != nil {
		return id, url, err
	}

	id = string(m["id"].(string))
	url = string(m["url"].(string))

	return id, url, nil
}
