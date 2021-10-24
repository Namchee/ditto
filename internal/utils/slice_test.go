package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSliceToString(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want string
	}{
		{
			name: "should join ints correctly",
			args: []int{1, 2, 3},
			want: "1, 2, 3",
		},
		{
			name: "should return empty string",
			args: []int{},
			want: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IntSliceToString(tc.args)

			assert.Equal(t, got, tc.want)
		})
	}
}
