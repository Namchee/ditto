package service

import "github.com/Namchee/ditto/internal/entity"

// Fetcher constructs the request
type Fetcher struct {
	endpoint *entity.Endpoint
}

// NewFetcher creates a new fetcher that can be used to do request calls
func NewFetcher(ep *entity.Endpoint) *Fetcher {
	return &Fetcher{endpoint: ep}
}

// Fetch constructs request and send it
func (f *Fetcher) Fetch() {

}
