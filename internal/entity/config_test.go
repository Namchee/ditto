package entity

import (
	"log"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		contents []byte
		want     *Configuration
	}{
		{
			name:     "should ignore when file does not exist",
			path:     "foo",
			contents: []byte{},
			want: &Configuration{
				TestDirectory: "ditto-test",
				LogDirectory:  "ditto-log",
				Strict:        false,
				Worker:        0,
				Status:        false,
				Parse:         false,
			},
		},
		{
			name:     "should ignore when file cannot be decoded",
			path:     "ditto.config.json",
			contents: []byte(`{ foo: "bar" }`),
			want: &Configuration{
				TestDirectory: "ditto-test",
				LogDirectory:  "ditto-log",
				Strict:        false,
				Worker:        0,
				Status:        false,
				Parse:         false,
			},
		},
		{
			name:     "should ignore invalid config",
			path:     "ditto.config.json",
			contents: []byte(`{ "test_directory": "bar", "log_directory": "baz", "strict": true, "worker": -1 }`),
			want: &Configuration{
				TestDirectory: "ditto-test",
				LogDirectory:  "ditto-log",
				Strict:        false,
				Worker:        0,
				Status:        false,
				Parse:         false,
			},
		},
		{
			name:     "should merge config",
			path:     "ditto.config.json",
			contents: []byte(`{ "test_directory": "bar", "log_directory": "baz", "strict": true, "worker": 2, "parse": true }`),
			want: &Configuration{
				TestDirectory: "bar",
				LogDirectory:  "baz",
				Strict:        true,
				Worker:        2,
				Status:        false,
				Parse:         true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fsys := fstest.MapFS{
				tc.path: {
					Data: tc.contents,
				},
			}
			got := ReadConfiguration(fsys, log.Default())

			assert.Equal(t, tc.want, got)
		})
	}
}
