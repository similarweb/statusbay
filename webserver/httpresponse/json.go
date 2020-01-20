package httpresponse

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// JSONWrite return JSON response to the client
func JSONWrite(resp http.ResponseWriter, statusCode int, data interface{}) error {

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	encoder := json.NewEncoder(resp)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// JSONError return JSON errors response
func JSONError(resp http.ResponseWriter, statusCode int, err error) error {
	return JSONWrite(resp, statusCode, HTTPErrorResponse{Error: err.Error()})
}

// JSONErrorParameters return JSON query validation
func JSONErrorParameters(resp http.ResponseWriter, statusCode int, err url.Values) error {
	e := map[string]interface{}{"Validation": err}
	return JSONWrite(resp, statusCode, e)
}
