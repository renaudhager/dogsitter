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

// List interface
type List interface {
	list(c *cli.Context) (err error)
	getDashboardList(ddEndpoint string, apiKey string, appKey string) (DashboardList, error)
	output(dashboardList DashboardList, format string, verbose bool) error
}

// ListAction struct
type ListAction struct{}

// NewListAction constructor for DeleteAction
func NewListAction() List {
	return &ListAction{}
}

// Dashboard structure definition, mapped from Datadog type
type Dashboard struct {
	CreatedAt   string `json:"created_at"`
	IsReadOnly  bool   `json:"is_read_only"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	LayoutType  string `json:"layout_type"`
	ModifiedAt  string `json:"modified_at"`
	Author      string `json:"author_handle"`
}

// DashboardList structure definition
type DashboardList struct {
	Dashboards []Dashboard `json:"dashboards"`
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
	Action: actionList,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format",
			Usage: "Format of the list of dashboard. Supported values are `text` or json.",
			Value: "text",
		},
		cli.StringFlag{
			Name:  "o, output",
			Usage: "output file to print dashboard list.",
			Value: "stdout",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "Get detail for a specific dashboard.",
		},
	},
}

// actionDelete placeholder function
func actionList(c *cli.Context) (err error) {
	return NewListAction().list(c)
}

func (la *ListAction) list(c *cli.Context) (err error) {

	var (
		dashboardList DashboardList
	)

	// Getting a list off all dashboards
	dashboardList, err = la.getDashboardList(c.GlobalString("dh"), c.GlobalString("api-key"), c.GlobalString("app-key"))

	if err != nil {
		log.Error("Error when retrieving dashboard list.")
		return err
	}

	id := c.String("id")

	if len(id) > 0 {
		var d DashboardList
		for _, dashboard := range dashboardList.Dashboards {
			if dashboard.ID == id {
				d.Dashboards = append(d.Dashboards, dashboard)
				la.output(d, c.String("format"), true)
			}
		}
	} else {
		la.output(dashboardList, c.String("format"), false)
	}

	return nil
}

// getDashboardList function that query Datadog to get list of dashboard
// then map the result into DashboardList structure.
func (la *ListAction) getDashboardList(ddEndpoint string, apiKey string, appKey string) (DashboardList, error) {

	var (
		dashboardList DashboardList
		body          []byte
		query         string
		statusCode    int
	)

	log.Info("Getting list of dashboards")

	query = ddEndpoint + "/api/v1/dashboard?api_key=" + apiKey + "&application_key=" + appKey

	resp, err := http.Get(query)

	if err != nil {
		log.Error("Error connection to ", query)
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

	err = json.Unmarshal(body, &dashboardList)

	if err != nil {
		log.Error("Error when unmarshalling Json to Struct.", err)
		return dashboardList, errors.New("Error when unmarshalling Json to Struct")
	}

	return dashboardList, nil
}

func (la *ListAction) output(dashboardList DashboardList, format string, verbose bool) error {

	switch format {
	case "text":
		if verbose {
			for _, dashboard := range dashboardList.Dashboards {
				fmt.Printf("    Title    |    Author    |    ID   |   Description   |    Created    |    Modified\n")
				fmt.Printf("%v | %v | %v| %v| %v| %v\n", dashboard.Title, dashboard.Author, dashboard.ID,
					dashboard.Description, dashboard.CreatedAt, dashboard.ModifiedAt)
			}
		} else {
			for _, dashboard := range dashboardList.Dashboards {
				fmt.Printf("%v | %v\n", dashboard.Title, dashboard.ID)
			}
		}
	case "json":
		b, _ := json.Marshal(dashboardList)
		fmt.Printf("%v\n", string(b))

	default:
		for _, dashboard := range dashboardList.Dashboards {
			fmt.Printf("%v | %v\n", dashboard.Title, dashboard.ID)
		}
	}

	return nil
}
