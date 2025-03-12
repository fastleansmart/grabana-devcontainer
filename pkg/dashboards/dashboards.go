package dashboards

import (
	"github.com/K-Phoen/grabana/dashboard"
)

type dashboardInitializer func() (dashboard.Builder, error)

// add dashboards to the array below
var dashboardInitializers = []dashboardInitializer{}

func InitDashboards() ([]dashboard.Builder, error) {
	registeredDashboards := make([]dashboard.Builder, len(dashboardInitializers))

	for i, init := range dashboardInitializers {
		d, err := init()
		if err != nil {
			return nil, err
		}

		registeredDashboards[i] = d
	}

	return registeredDashboards, nil
}
