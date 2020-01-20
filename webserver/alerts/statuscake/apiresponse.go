package statuscake

// Test represents the JSON response from Statuscake API.
type Test struct {
	TestID      int    `json:"TestID" querystring:"TestID" querystringoptions:"omitempty"`
	WebsiteName string `json:"WebsiteName" querystring:"WebsiteName"`
}

// Periods represents the JSON response from Statuscake API.
type Periods struct {
	Status     string `json:"Status" querystring:"Status"`
	Start      string `json:"Start" querystring:"Start"`
	End        string `json:"End" querystring:"End"`
	StartUnix  int64  `json:"Start_Unix" querystring:"Start_Unix"`
	EndUnix    int64  `json:"End_Unix" querystring:"End_Unix"`
	Additional string `json:"Additional" querystring:"Additional"`
	Period     string `json:"Period" querystring:"Period"`
}
