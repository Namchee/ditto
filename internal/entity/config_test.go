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
				Directory: "ditto-test",
				Strict:    false,
			},
		},
		{
			name:     "should ignore when file cannot be decoded",
			path:     "ditto.config.json",
			contents: []byte(`{ foo: "bar" }`),
			want: &Configuration{
				Directory: "ditto-test",
				Strict:    false,
			},
		},
		{
			name:     "should merge config",
			path:     "ditto.config.json",
			contents: []byte(`{ "directory": "bar", "strict": true }`),
			want: &Configuration{
				Directory: "bar",
				Strict:    true,
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
