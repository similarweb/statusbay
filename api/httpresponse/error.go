package httpresponse

// HTTPErrorResponse is returned on error
type HTTPErrorResponse struct {
	Error string `json:"error"`
}
