package utils

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestIsFileExists(t *testing.T) {
	tests := []struct {
		name string
		args string
		mock fstest.MapFS
		want bool
	}{
		{
			name: "file exist",
			args: "foo.txt",
			mock: fstest.MapFS{
				"foo.txt": {
					Data: []byte("hello"),
				},
			},
			want: true,
		},
		{
			name: "file not exist",
			args: "foo.txt",
			mock: fstest.MapFS{
				"bar.txt": {
					Data: []byte("hello"),
				},
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsFileExist(tc.mock, tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
