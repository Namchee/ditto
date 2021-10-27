package service

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/stretchr/testify/assert"
)

var (
	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		q := req.URL.Query()

		if q.Has("err") {
			// Simulate faulty response body
			rw.Header().Set("Content-Length", "1")
			return
		}

		var err error

		if q.Has("foo") {
			_, err = rw.Write([]byte(`{ "bar": "baz" }`))
		} else {
			_, err = rw.Write([]byte(`{ "foo": "bar" }`))
		}

		if err != nil {
			panic("shouldn't happen")
		}
	}))
)

func TestFetcher_Fetch(t *testing.T) {

	tests := []struct {
		name     string
		endpoint entity.Endpoint
		want     *entity.FetchResult
		err      error
	}{
		{
			name: "should throw an error when request cannot be created",
			endpoint: entity.Endpoint{
				Method: "(",
				Host:   server.URL,
			},
			want: nil,
			err:  constant.ErrCreateRequest,
		},
		{
			name: "should throw an error when request cannot be sent",
			endpoint: entity.Endpoint{
				Method: "GET",
			},
			want: nil,
			err:  errors.New(`[Fetcher] Failed to fetch response: Get "": unsupported protocol scheme ""`),
		},
		{
			name: "should throw an error when response body is malformed",
			endpoint: entity.Endpoint{
				Method: "GET",
				Host:   server.URL,
				Query: map[string]interface{}{
					"err": 1,
				},
			},
			want: nil,
			err:  constant.ErrReadResponse,
		},
		{
			name: "success with plain body",
			endpoint: entity.Endpoint{
				Method: "GET",
				Host:   server.URL,
			},
			want: &entity.FetchResult{
				Endpoint: entity.Endpoint{
					Method: "GET",
					Host:   server.URL,
				},
				Status:   200,
				Response: `{ "foo": "bar" }`,
			},
			err: nil,
		},
		{
			name: "success with query",
			endpoint: entity.Endpoint{
				Method: "GET",
				Host:   server.URL,
				Query: map[string]interface{}{
					"foo": "bar",
				},
			},
			want: &entity.FetchResult{
				Endpoint: entity.Endpoint{
					Method: "GET",
					Host:   server.URL,
					Query: map[string]interface{}{
						"foo": "bar",
					},
				},
				Status:   200,
				Response: `{ "bar": "baz" }`,
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := NewFetcher(tc.endpoint)

			got, err := f.Fetch()

			fmt.Println(err)
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.err, err)
		})
	}
}
