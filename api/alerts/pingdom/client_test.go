package pingdom

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"runtime"
	"testing"

	"statusbay/request"
)

type MockHTTPClient struct {
	RequestData *http.Request

	ReturnStatusCode int
	ReturnBodyPath   string
	ReturnError      error
}

func (mh *MockHTTPClient) Request(method string, url string, v url.Values, body io.Reader) (*http.Request, error) {

	req := request.NewHTTPClient()

	return req.Request(method, url, v, nil)
}
func (mh *MockHTTPClient) DO(r *http.Request) (*http.Response, error) {

	if mh.ReturnError != nil {
		return nil, mh.ReturnError
	}
	mh.RequestData = r

	response := http.Response{}
	response.StatusCode = mh.ReturnStatusCode

	_, filename, _, _ := runtime.Caller(0)
	currentFolderPath := filepath.Dir(filename)
	reader, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", currentFolderPath, mh.ReturnBodyPath))

	response.Body = ioutil.NopCloser(bytes.NewBufferString(string(reader)))
	return &response, nil
}

func TestNewRequest(t *testing.T) {

	testsCase := []struct {
		name         string
		bodyFilePath string
		token        string
		urlValues    url.Values

		expectedURL        string
		expectedStatusCode int
		expectedToken      string
		expectedErr        error
	}{
		{"Valid response", "testutils/mock/checks.json", "foo1234", nil, "127.0.0.1/3.1/path", http.StatusOK, "Bearer foo1234", nil},
		{"URL parameters", "testutils/mock/checks.json", "foo1234", url.Values{"foo": []string{"foo"}}, "127.0.0.1/3.1/path?foo=foo", http.StatusOK, "Bearer foo1234", nil},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: test.expectedStatusCode,
				ReturnBodyPath:   test.bodyFilePath,
			}

			client := NewClient("127.0.0.1", test.token, mockHTTP)
			response, err := client.newRequest("GET", "path", test.urlValues, nil)

			if err != test.expectedErr {
				t.Fatalf("unexpected response, got %d, expected %d", err, test.expectedErr)
			}

			if mockHTTP.RequestData.URL.String() != test.expectedURL {
				t.Fatalf("unexpected URL, got %s, expected %s", mockHTTP.RequestData.URL, test.expectedURL)
			}

			if mockHTTP.RequestData.Header["Authorization"][0] != test.expectedToken {
				t.Fatalf("unexpected URL, expected %s", test.expectedToken)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Fatalf("unexpected status code, got %d, expected %d", response.StatusCode, test.expectedStatusCode)
			}

		})
	}

}

func TestGetCheckSummaryOutage(t *testing.T) {

	testsCase := []struct {
		name               string
		bodyFilePath       string
		err                error
		expectedCheckCount int
	}{
		{"Valid response", "testutils/mock/summary.outage.json", nil, 5},
		{"Valid response", "testutils/mock/checks.json", errors.New("error"), 3},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: http.StatusOK,
				ReturnBodyPath:   test.bodyFilePath,
				ReturnError:      test.err,
			}

			client := NewClient("127.0.0.1", "", mockHTTP)

			resp, err := client.GetCheckSummaryOutage(1, nil)

			if err != nil {
				if err != test.err {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}

			if len(resp.Summary.States) != test.expectedCheckCount {
				t.Fatalf("unexpected metrics length response, got %d expected %d", len(resp.Summary.States), test.expectedCheckCount)
			}

		})
	}

}

func TestGetChecks(t *testing.T) {

	testsCase := []struct {
		name               string
		bodyFilePath       string
		err                error
		expectedCheckCount int
	}{
		{"Valid response", "testutils/mock/checks.json", nil, 3},
		{"Valid response", "testutils/mock/checks.json", errors.New("error"), 3},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: http.StatusOK,
				ReturnBodyPath:   test.bodyFilePath,
				ReturnError:      test.err,
			}

			client := NewClient("127.0.0.1", "", mockHTTP)

			resp, err := client.GetChecks(nil)

			if err != nil {
				if err != test.err {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}

			if len(resp.Checks) != test.expectedCheckCount {
				t.Fatalf("unexpected tests count response, got %d expected %d", len(resp.Checks), test.expectedCheckCount)
			}

		})
	}

}
