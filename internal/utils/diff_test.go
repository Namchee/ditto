package utils

import (
	"testing"

	"github.com/Namchee/ditto/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestHasDiff(t *testing.T) {
	tests := []struct {
		name   string
		args   []*entity.FetchResult
		config bool
		want   bool
	}{
		{
			name: "should return false",
			args: []*entity.FetchResult{
				{
					Status:   200,
					Response: "foo",
				},
				{
					Status:   200,
					Response: "foo",
				},
				{
					Status:   200,
					Response: "foo",
				},
			},
			config: true,
			want:   false,
		},
		{
			name: "should return true when diff",
			args: []*entity.FetchResult{
				{
					Status:   200,
					Response: "foo",
				},
				{
					Status:   200,
					Response: "bar",
				},
				{
					Status:   200,
					Response: "baz",
				},
			},
			config: true,
			want:   true,
		},
		{
			name: "should return false when status diff but not strict",
			args: []*entity.FetchResult{
				{
					Status:   200,
					Response: "foo",
				},
				{
					Status:   301,
					Response: "foo",
				},
				{
					Status:   200,
					Response: "foo",
				},
			},
			config: false,
			want:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			config := &entity.Configuration{
				Status: tc.config,
			}
			got := HasDiff(tc.args, config)

			assert.Equal(t, tc.want, got)
		})
	}
}
