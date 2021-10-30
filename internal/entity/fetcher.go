package entity

// FetchResult represents fetching result when testing endpoints
type FetchResult struct {
	Endpoint
	Status   int    `json:"status"`
	Response string `json:"response"`
}
