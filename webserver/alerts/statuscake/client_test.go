package statuscake

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
	http.NewRequest(method, url, nil)
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
		urlValues    url.Values

		expectedUsername   string
		expectedKey        string
		expectedURL        string
		expectedStatusCode int
		expectedErr        error
	}{
		{"Valid response", "testutils/mock/tests.json", nil, "username", "key", "127.0.0.1/path", http.StatusOK, nil},
		{"URL parameters", "testutils/mock/tests.json", url.Values{"foo": []string{"foo"}}, "username", "key", "127.0.0.1/path?foo=foo", http.StatusOK, nil},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: test.expectedStatusCode,
				ReturnBodyPath:   test.bodyFilePath,
			}

			client := NewClient("127.0.0.1", "username", "key", mockHTTP)
			response, err := client.newRequest("GET", "path", test.urlValues, nil)

			if err != test.expectedErr {
				t.Fatalf("unexpected error response, got %d, expected %d", err, test.expectedErr)
			}

			if mockHTTP.RequestData.URL.String() != test.expectedURL {
				t.Fatalf("unexpected http URL, got %s, expected %s", mockHTTP.RequestData.URL, test.expectedURL)
			}

			if mockHTTP.RequestData.Header["Username"][0] != test.expectedUsername {
				t.Fatalf("unexpected http URL, got %s expected %s", mockHTTP.RequestData.Header["Username"][0], test.expectedUsername)
			}

			if mockHTTP.RequestData.Header["Api"][0] != test.expectedKey {
				t.Fatalf("unexpected http URL, got %s expected %s", mockHTTP.RequestData.Header["Api"][0], test.expectedKey)
			}

			if response.StatusCode != test.expectedStatusCode {
				t.Fatalf("unexpected http status code, got %d, expected %d", response.StatusCode, test.expectedStatusCode)
			}

		})
	}

}

func TestGetTests(t *testing.T) {

	testsCase := []struct {
		name               string
		bodyFilePath       string
		err                error
		expectedCheckCount int
	}{
		{"Valid response", "testutils/mock/tests.json", nil, 2},
		{"Valid response", "testutils/mock/tests.json", errors.New("error"), 3},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: http.StatusOK,
				ReturnBodyPath:   test.bodyFilePath,
				ReturnError:      test.err,
			}

			client := NewClient("127.0.0.1", "username", "password", mockHTTP)

			resp, err := client.GetTests(nil)

			if err != nil {
				if err != test.err {
					t.Fatalf("unexpected error response")
				}
				return
			}

			if len(resp) != test.expectedCheckCount {
				t.Fatalf("unexpected test count response, got %d expected %d", len(resp), test.expectedCheckCount)
			}

		})
	}

}

func TestPeriods(t *testing.T) {

	testsCase := []struct {
		name               string
		bodyFilePath       string
		err                error
		expectedCheckCount int
	}{
		{"Valid response", "testutils/mock/tests.periods.json", nil, 3},
		{"Valid response", "testutils/mock/tests.periods.json", errors.New("error"), 3},
	}

	for _, test := range testsCase {
		t.Run(test.name, func(t *testing.T) {

			mockHTTP := &MockHTTPClient{
				ReturnStatusCode: http.StatusOK,
				ReturnBodyPath:   test.bodyFilePath,
				ReturnError:      test.err,
			}

			client := NewClient("127.0.0.1", "username", "password", mockHTTP)

			resp, err := client.GetTests(nil)

			if err != nil {
				if err != test.err {
					t.Fatalf("unexpected error response")
				}

				return
			}

			if len(resp) != test.expectedCheckCount {
				t.Fatalf("unexpected test count response, got %d expected %d", len(resp), test.expectedCheckCount)
			}

		})
	}

}
