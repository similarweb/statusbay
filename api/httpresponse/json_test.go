package httpresponse_test

import (
	"errors"
	"net/http"
	"statusbay/api/httpresponse"
	"testing"
)

type FakeResponse struct {
	t       *testing.T
	headers http.Header
	body    []byte
	status  int
}

func NewFaceResponse(t *testing.T) *FakeResponse {
	return &FakeResponse{
		t:       t,
		headers: make(http.Header),
	}
}

func (r *FakeResponse) Header() http.Header {
	return r.headers
}

func (r *FakeResponse) Write(body []byte) (int, error) {
	r.body = body
	return len(body), nil
}

func (r *FakeResponse) WriteHeader(status int) {
	r.status = status
}

func TestJSONWrite(t *testing.T) {

	ts := NewFaceResponse(t)
	httpresponse.JSONWrite(ts, 200, `{"key": "value"}`)

	if len(ts.headers) != 1 {
		t.Fatalf("unexpected json header counts, got %d expected %d", len(ts.headers), 1)
	}

	if ts.status != 200 {
		t.Fatalf("unexpected json status code response, got %d expected %d", ts.status, 200)
	}

}

func TestJSONError(t *testing.T) {

	ts := NewFaceResponse(t)
	err := errors.New("error message")
	httpresponse.JSONError(ts, 201, err)

	expected := `{
  "error": "error message"
}
`
	if string(ts.body) != expected {
		t.Fatalf("unexpected json header counts, got %s expected %s", string(ts.body), expected)
	}

}
