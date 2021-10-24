package service

import (
	"errors"
	"testing"

	"github.com/Namchee/ditto/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	tests := []struct {
		name string
		args *entity.TestResult
		want string
	}{
		{
			name: "should format error test",
			args: &entity.TestResult{
				Name:  "TestThisOne",
				Error: errors.New("foo bar"),
				Diff:  []int{},
			},
			want: "TestThisOne: ❌ FAIL = Failed to run test: foo bar",
		},
		{
			name: "should format passed test",
			args: &entity.TestResult{
				Name: "TestThisOne",
				Diff: []int{},
			},
			want: "TestThisOne: ✔️ PASS",
		},
		{
			name: "should format failed test",
			args: &entity.TestResult{
				Name: "TestThisOne",
				Diff: []int{1},
			},
			want: "TestThisOne: ❌ FAIL = Endpoint(s) with index 1 have different result(s)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatResult(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}
