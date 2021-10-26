package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

// Fetcher constructs the request
type Fetcher struct {
	endpoint entity.Endpoint
	client   *http.Client
}

// NewFetcher creates a new fetcher that can be used to do request calls
func NewFetcher(ep entity.Endpoint) *Fetcher {
	return &Fetcher{
		endpoint: ep,
		client: &http.Client{
			Timeout: time.Duration(ep.Timeout),
		},
	}
}

// Fetch constructs request, send the request, and return a response string from it
func (f *Fetcher) Fetch() (string, error) {
	query := url.Values{}

	for k, v := range f.endpoint.Query {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	for k, v := range f.endpoint.Body {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	request, err := http.NewRequest(f.endpoint.Method, f.endpoint.Host, nil)

	if err != nil {
		return "", constant.ErrCreateRequest
	}

	request.URL.RawQuery = query.Encode()

	for k, v := range f.endpoint.Headers {
		request.Header.Add(k, v)
	}
	request.Header.Add("Content-Length", strconv.Itoa(len(query.Encode())))

	resp, err := f.client.Do(request)

	if err != nil {
		return "", constant.ErrFetchResponse
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", constant.ErrReadResponse
	}

	return string(body), nil
}
