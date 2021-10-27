package utils

import (
	"errors"
	"testing"

	"github.com/Namchee/ditto/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	tests := []struct {
		name   string
		args   *entity.RunnerResult
		status bool
		want   string
	}{
		{
			name: "should format error test",
			args: &entity.RunnerResult{
				Name:   "TestThisOne",
				Error:  errors.New("foo bar"),
				Result: []*entity.FetchResult{},
			},
			status: false,
			want:   "TestThisOne: ❌ FAIL = Failed to run test: foo bar",
		},
		{
			name: "should format passed test",
			args: &entity.RunnerResult{
				Name:   "TestThisOne",
				Result: []*entity.FetchResult{},
			},
			status: true,
			want:   "TestThisOne: ✅ PASS",
		},
		{
			name: "should format failed test",
			args: &entity.RunnerResult{
				Name:   "TestThisOne",
				Result: []*entity.FetchResult{},
			},
			status: false,
			want:   "TestThisOne: ❌ FAIL. Please check the generated test log.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatResult(tc.args, tc.status)

			assert.Equal(t, tc.want, got)
		})
	}
}
