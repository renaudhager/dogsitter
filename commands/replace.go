package commands

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Replace interface
type Replace interface {
	replaceDashboard(c *cli.Context) error
}

// ReplaceAction struct
type ReplaceAction struct{}

// NewReplaceAction constructor for ReplaceAction
func NewReplaceAction() Replace {
	return &ReplaceAction{}
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// ReplaceCmd pull to download dashboard configuration from Datadog.
var ReplaceCmd = cli.Command{
	Name:   "replace",
	Usage:  "Replace dashboard from Datadog.",
	Action: actionReplace,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "title",
			Usage: "Dashboard title to replace. Cannot be used with --id.",
		},
		cli.StringFlag{
			Name:  "id",
			Usage: "Dashboard id to replace. Cannot be used with --title.",
		},
		cli.StringFlag{
			Name:  "f, file",
			Usage: "Dashboard file configuration.",
		},
	},
}

// actionReplace placeholder function
func actionReplace(c *cli.Context) (err error) {
	return NewReplaceAction().replaceDashboard(c)
}

func (da *ReplaceAction) replaceDashboard(c *cli.Context) error {

	dashboardID, err := searchDashboard(c.GlobalString("dh"), c.GlobalString("api-key"), c.GlobalString("app-key"), c.String("title"))

	if err != nil {
		log.Errorf("Error findind the dashboard: %v", err)
		return err
	}

	log.Infof("ID: %s", dashboardID)

	return nil
}

// searchDashboard this function return the ID of the 1st dashboard that matches the title
func searchDashboard(ddEndpoint string, apiKey string, appKey string, title string) (string, error) {

	var (
		dashboardList   DashboardList
		wantedDashboard Dashboard
	)

	la := NewListAction()
	dashboardList, err := la.getDashboardList(ddEndpoint, apiKey, appKey)

	if err != nil {
		return "", err
	}

	for _, dashboard := range dashboardList.Dashboards {
		if dashboard.Title == title {
			wantedDashboard = dashboard
		}
	}

	if wantedDashboard.ID != "" {
		return wantedDashboard.ID, nil
	}

	return "", errors.New("Dashboard not found")
}
