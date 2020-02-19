package statuscake

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"statusbay/request"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// ClientDescriber descrive the statuscake client available function
type ClientDescriber interface {
	GetTests(filterOptions url.Values) ([]*Test, error)
	Periods(ID int) ([]*Periods, error)
}

type responseBody struct {
	io.Reader
}

func (r *responseBody) Close() error {
	return nil
}

// authErrorResponse represents a authentication error details
type authErrorResponse struct {
	ErrNo int
	Error string
}

// Client struct
type Client struct {
	client   request.HTTPClientDescriber
	endpoint string
	username string
	key      string
}

// NewClient creates new statuscake client
func NewClient(endpoint, username, key string, client request.HTTPClientDescriber) *Client {
	return &Client{
		client:   client,
		endpoint: endpoint,
		username: username,
		key:      key,
	}
}

// newRequest Prepare the header parameters and execute the request
func (sc *Client) newRequest(method string, path string, v url.Values, body io.Reader) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", sc.endpoint, path)

	lg := log.WithFields(log.Fields{
		"method": method,
		"path":   path,
		"values": v,
	})

	req, err := sc.client.Request("GET", url, v, body)

	if err != nil {
		lg.WithError(err).Error("could not create HTTP client request")
		return nil, err
	}

	req.Header.Set("Username", sc.username)
	req.Header.Set("API", sc.key)

	resp, err := sc.client.DO(req)
	if err != nil {
		lg.WithError(err).Error("could not send HTTP client request")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		lg.WithField("status_code", resp.StatusCode).Error("unexpected status code returned")
		return nil, &request.HttpError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}
	}

	var authError authErrorResponse

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &authError)
	if err == nil && authError.ErrNo == 0 && authError.Error != "" {
		return nil, errors.New(fmt.Sprintf("%d, %s", authError.ErrNo, authError.Error))
	}

	resp.Body = &responseBody{
		Reader: bytes.NewReader(b),
	}

	return resp, nil

}

// GetTests returns list of tests with a given filter options
func (sc *Client) GetTests(filterOptions url.Values) ([]*Test, error) {

	resp, err := sc.newRequest("GET", "/Tests", filterOptions, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tests []*Test
	err = json.NewDecoder(resp.Body).Decode(&tests)

	return tests, err

}

// Periods return period check status from a given check ID and filter options
func (sc *Client) Periods(ID int) ([]*Periods, error) {
	v := url.Values{}
	v.Set("TestID", strconv.Itoa(ID))
	resp, err := sc.newRequest("GET", "/Tests/Periods", v, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tests []*Periods
	err = json.NewDecoder(resp.Body).Decode(&tests)

	return tests, err
}
