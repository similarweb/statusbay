package pingdom

// ChecksResponse represents the JSON response from checks Pingdom API.
type ChecksResponse struct {
	Checks []Check `json:"checks"`
}

// Check represents the JSON response for a check from the Pingdom API.
type Check struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

// SummaryOutageResponse represents the JSON response from summary.outage Pingdom API.
type SummaryOutageResponse struct {
	Summary States `json:"summary"`
}

// States represents the JSON response from summary.outage Pingdom API.
type States struct {
	States []State `json:"states"`
}

// State represents the JSON response the state check.
type State struct {
	Status   string `json:"status"`
	TimeFrom int64  `json:"timefrom"`
	TimeTo   int64  `json:"timeto"`
}
