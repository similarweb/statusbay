package statuscake

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

type MockStatuscakeClient struct {
	returnChecksResponse  []*Test
	returnPeriodsResponse map[int][]*Periods

	getChecksFilters url.Values
}

func (mp *MockStatuscakeClient) GetTests(filterOptions url.Values) ([]*Test, error) {

	mp.getChecksFilters = filterOptions
	return mp.returnChecksResponse, nil
}

func (mp *MockStatuscakeClient) Periods(ID int) ([]*Periods, error) {

	response := mp.returnPeriodsResponse[ID]
	return response, nil
}

func MockStatuscakeManager(checkResponse []*Test, summaryOutageResponse map[int][]*Periods) (*Statuscake, *MockStatuscakeClient) {
	mockClient := &MockStatuscakeClient{
		returnChecksResponse:  checkResponse,
		returnPeriodsResponse: summaryOutageResponse,
	}

	statuscakeManager := NewStatuscakeManager(mockClient)
	return statuscakeManager, mockClient
}

func TestGetAlertByTags(t *testing.T) {

	checkResponse := []*Test{
		{TestID: 1, WebsiteName: "foo"},
		{TestID: 2, WebsiteName: "foo2"},
	}

	periodsResponse := map[int][]*Periods{
		1: []*Periods{
			&Periods{Status: "up"},
		},
		2: []*Periods{
			{Status: "up", StartUnix: 1577836801, EndUnix: 1577923100},
			{Status: "down", StartUnix: 1577836802, EndUnix: 1577923100},
			{Status: "down", StartUnix: 1, EndUnix: 2},
		},
	}

	statuscakeManager, mockClient := MockStatuscakeManager(checkResponse, periodsResponse)

	from := time.Date(2020, time.January, 01, 0, 0, 0, 0, time.UTC)
	to := time.Date(2020, time.January, 02, 0, 0, 0, 0, time.UTC)

	fmt.Println(from.Unix())
	fmt.Println(to.Unix())
	checks, err := statuscakeManager.GetAlertByTags("foo,foo2", from, to)

	t.Run("validate checks", func(t *testing.T) {

		if err != nil {
			t.Fatalf("unexpected error")
		}

		if len(checks) != 2 {
			t.Fatalf("unexpected checks count, got %d expected %d", len(checks), 2)
		}

		check := checks[0]

		if check.ID != 2 {
			t.Fatalf("unexpected check ID, got %d expected %d", check.ID, 2)
		}

		if check.URL != fmt.Sprintf(PageURL, check.ID) {
			t.Fatalf("unexpected check URL, got %s expected %s", check.URL, fmt.Sprintf(PageURL, check.ID))
		}

		if check.Name != "foo2" {
			t.Fatalf("unexpected check Name, got %s expected %s", check.Name, "foo2")
		}

		if len(check.Periods) != 2 {
			t.Fatalf("unexpected periods check count, got %d expected %d", len(check.Periods), 2)
		}

		if check.Periods[0].Status != "up" {
			t.Fatalf("unexpected periods check count, got %s expected %s", check.Periods[0].Status, "up")
		}

		if check.Periods[0].StartUnix != 1577836801 {
			t.Fatalf("unexpected periods check count, got %d expected %d", check.Periods[0].StartUnix, 1577836801)
		}

		if check.Periods[0].EndUnix != 1577923100 {
			t.Fatalf("unexpected periods check count, got %d expected %d", check.Periods[0].EndUnix, 1577923100)
		}

	})

	t.Run("validate checks request queries", func(t *testing.T) {
		if len(mockClient.getChecksFilters["tags"]) != 1 {
			t.Fatalf("unexpected tags filter query, got %d expected %d", len(mockClient.getChecksFilters["tags"]), 1)
		}

		if mockClient.getChecksFilters["tags"][0] != "foo,foo2" {
			t.Fatalf("unexpected tags filter query value, got %s expected %s", mockClient.getChecksFilters["tags"][0], "foo,foo2")
		}

	})

}
