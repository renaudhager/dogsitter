package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Dashboard structure definition, mapped from Datadog type
type Dashboard struct {
	createdAt   string `json:"created_at"`
	isReadOnly  bool   `json:"is_read_only"`
	description string `json:  "description"`
	ID          string `json: "id"`
	Title       string `json:  "title"`
	URL         string `json: "url"`
	layoutType  string `json:  "layout_type"`
	modifiedAt  string `json: "modified_at"`
	Author      string `json: "author_handle"`
}

// DashboardList structure definition
type DashboardList struct {
	dashboards []Dashboard `json:"dashboards"`
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// ListCmd pull to List existing dashboard defined in Datadog.
var ListCmd = cli.Command{
	Name:   "list",
	Usage:  "List dashboard existing in Datadog.",
	Action: list,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "Format of the list of dashboard.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "o, output",
			Usage: "output file to print dashboard list.",
			Value: "stdout",
		},
	},
}

func list(c *cli.Context) (err error) {

	var (
	// dashboardList DashboardList
	)

	_, err = getDashboardList(c.GlobalString("dh"), c.GlobalString("api-key"), c.GlobalString("app-key"))
	return nil
}

func getDashboardList(ddEndpoint string, apiKey string, appKey string) (DashboardList, error) {

	var (
		dashboardList DashboardList
		body          []byte
		query         string
		statusCode    int
	)

	log.Info("Getting list of dashboards")

	query = ddEndpoint + "/api/v1/dashboard?api_key=" + apiKey + "&application_key=" + appKey
	fmt.Printf("query: %v \n", query)
	resp, err := http.Get(query)

	if err != nil {
		log.Error("Error connectiong to ", query)
		return dashboardList, err
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode

	if statusCode != 200 {
		log.Error("Returned code is not 200, it's ", statusCode)
		return dashboardList, errors.New("Returned code is not 200")
	}

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Error("Error when reading body response. ", err)
		return dashboardList, errors.New("Error when reading body response")
	}

	fmt.Printf("\nbody:\n%v", string(body))

	err = json.Unmarshal(body, &dashboardList)

	if err != nil {
		log.Error("Error when unmarshalling Json to Struct.", err)
		return dashboardList, errors.New("Error when unmarshalling Json to Struct")
	}

	fmt.Printf("\nlist: %v", dashboardList)

	return dashboardList, nil
}
