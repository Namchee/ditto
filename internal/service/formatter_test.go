package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatResult(t *testing.T) {
	type args struct {
		test string
		diff []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should format passed test",
			args: args{
				test: "TestThisOne",
				diff: []int{},
			},
			want: "TestThisOne: ✔️ PASS",
		},
		{
			name: "should format failed test",
			args: args{
				test: "TestThisOne",
				diff: []int{1},
			},
			want: "TestThisOne: ❌ FAIL = Endpoint(s) with index 1 have different result(s)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatResult(tc.args.test, tc.args.diff)

			assert.Equal(t, tc.want, got)
		})
	}
}
