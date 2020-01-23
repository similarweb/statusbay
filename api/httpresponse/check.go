package httpresponse

//HealthResponse is returned when healtcheck requested
type HealthResponse struct {
	Status bool `json:"status"`
}
