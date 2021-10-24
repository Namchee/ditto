package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDiff(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []int
	}{
		{
			name: "should return empty array",
			args: []string{"foo", "foo", "foo"},
			want: []int{},
		},
		{
			name: "should return diff list",
			args: []string{"foo", "bar", "baz"},
			want: []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetDiff(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
