package entity

// FetchResult represents fetching result when testing endpoints
type FetchResult struct {
	Endpoint
	Status   int
	Response string
}
