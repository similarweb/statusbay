package httpresponse

type CheckResponse struct {
	ID      int
	URL     string
	Name    string
	Periods []PeriodsResponse
}

type PeriodsResponse struct {
	Status    string
	StartUnix int64
	EndUnix   int64
}
