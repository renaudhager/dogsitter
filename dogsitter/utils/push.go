package utils

import (
	"bytes"
	"errors"
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

	// fmt.Println(content)

	// c.GlobalString("dh")
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

	url := endpoint + "/api/v2/dashboard?api_key=" + apiKey + "&application_key=" + appKey
	contentIO := bytes.NewBuffer([]byte(content))

	req, err := http.NewRequest("POST", url, contentIO)
	if err != nil {
		log.Error("Error during POST request creation: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error executing POST request:", err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode

	if statusCode != 200 {
		log.Error("Error returned is not 200:", statusCode)
		err = errors.New("Error during dashboard import")
		return err
	}

	log.Info("Dashboard successfully imported.")

	return nil
}
