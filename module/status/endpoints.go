package status

import whitebit "github.com/whitebit-exchange/go-sdk"

const maintenanceStatusUrl = "/api/v4/public/platform/status"

type maintenanceStatus struct {
	whitebit.NoAuth
}

func newMaintenanceStatusEndpoint() *maintenanceStatus {
	return &maintenanceStatus{}
}

func (m *maintenanceStatus) Url() string {
	return maintenanceStatusUrl
}
