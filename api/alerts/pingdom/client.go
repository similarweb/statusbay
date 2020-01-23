package pingdom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"statusbay/request"

	log "github.com/sirupsen/logrus"
)

const (
	// APIVersion present the API version
	APIVersion = "3.1"
)

// ClientDescriber describe the Pingdom client available function
type ClientDescriber interface {
	GetChecks(filterOptions url.Values) (*ChecksResponse, error)
	GetCheckSummaryOutage(ID int, filterOptions url.Values) (*SummaryOutageResponse, error)
}

type responseBody struct {
	io.Reader
}

func (r *responseBody) Close() error {
	return nil
}

// Client struct
type Client struct {
	client   request.HTTPClientDescriber
	endpoint string
	token    string
}

// NewClient creates new Pingdom client
func NewClient(endpoint, token string, client request.HTTPClientDescriber) *Client {
	return &Client{
		client:   client,
		endpoint: endpoint,
		token:    token,
	}
}

// newRequest prepare the http Pingdom request
func (pc *Client) newRequest(method string, path string, v url.Values, body io.Reader) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s/%s", pc.endpoint, APIVersion, path)

	lg := log.WithFields(log.Fields{
		"method": method,
		"url":    url,
		"values": v.Encode(),
	})

	req, err := pc.client.Request("GET", url, v, body)

	if err != nil {
		lg.WithError(err).Error("Error when trying to create HTTP client request")
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", pc.token))

	resp, err := pc.client.DO(req)
	if err != nil {
		lg.WithError(err).Error("Error when trying to send HTTP request")
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		lg.WithError(err).Error("HTTP request failed")
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		lg.WithField("status_code", resp.StatusCode).Error("HTTP response return with invalid status code")
		return nil, &request.HttpError{
			Status:     resp.Status,
			StatusCode: resp.StatusCode,
		}
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		lg.WithError(err).Error("Error when trying to read http body response")
		return nil, err
	}

	resp.Body = &responseBody{
		Reader: bytes.NewReader(b),
	}

	return resp, nil

}

// GetChecks return list of Pingdom checks
func (pc *Client) GetChecks(filterOptions url.Values) (*ChecksResponse, error) {

	resp, err := pc.newRequest("GET", "/checks", filterOptions, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var checks *ChecksResponse

	err = json.NewDecoder(resp.Body).Decode(&checks)

	return checks, err

}

// GetCheckSummaryOutage return summary outage from a given check ID and filter options
func (pc *Client) GetCheckSummaryOutage(ID int, filterOptions url.Values) (*SummaryOutageResponse, error) {

	resp, err := pc.newRequest("GET", fmt.Sprintf("/summary.outage/%d", ID), filterOptions, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var results *SummaryOutageResponse

	err = json.NewDecoder(resp.Body).Decode(&results)

	return results, err

}
